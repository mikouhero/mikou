package app

import "mikou/global"

// 通用开启协程方法
func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				global.LoggerV2.Errorf("go func err: %s", err)
				// todo notice
			}
		}()
		f()
	}()
}
