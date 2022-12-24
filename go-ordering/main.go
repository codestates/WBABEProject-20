package main

import (
	conf "WBABEProject-20/go-ordering/conf"
	ctl "WBABEProject-20/go-ordering/controller"
	"WBABEProject-20/go-ordering/logger"
	"WBABEProject-20/go-ordering/model"
	rt "WBABEProject-20/go-ordering/router"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @title WBA [Backend Final Project]
// @version 1.0
// @description 띵동주문이요, 온라인 주문 시스템(Online Ordering System)
func main() {

	//conf
	var configFlag = flag.String("config", "./conf/config.toml", "toml file to use for configuration")
	flag.Parse()
	cf := conf.GetConfig(*configFlag)

	// 로그 초기화
	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	logger.Debug("ready server....")

	//model 모듈 선언
	if mod, err := model.NewModel(cf); err != nil {
		logger.Error(err)
		panic(err)
	} else if controller, err := ctl.NewCTL(mod); err != nil { //controller 모듈 설정
		logger.Error(err)
		panic(err)
	} else if rt, err := rt.NewRouter(controller); err != nil { //router 모듈 설정
		logger.Error(err)
		panic(err)
	} else {
		fmt.Println("main else ")
		mapi := &http.Server{
			Addr:           cf.Server.Port,
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		go func() {
			if err := mapi.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		stopSig := make(chan os.Signal) //chan 선언
		// 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
		signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM)
		<-stopSig //메세지 등록
		logger.Warn("Shutdown Server ...")

		// 해당 context 타임아웃 설정, 5초후 server stop
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			logger.Error("Server Shutdown:", err)
		}
		// catching ctx.Done(). timeout of 5 seconds.
		select {
		case <-ctx.Done():
			logger.Info("timeout of 5 seconds.")
		}
		logger.Info("Server exiting")

	}

	if err := g.Wait(); err != nil {
		logger.Error(err)
	}

}
