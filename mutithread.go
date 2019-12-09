package main

import ("fmt"
        "time")

func counter(out1 chan<-int,out2 chan<-int,out3 chan<-int) {
	for x :=0;x<100;x++{
		time.Sleep(200*time.Millisecond)
		//依次将x放入对应的channel中
		if x%3==0 {
			out1<-x;
		}else if x%3==1{
			out2<-x;
		}else{
			out3<-x;
		}
		
	}
	close(out1)
	close(out2)
	close(out3)
}

func squarer(out1 chan<-int,out2 chan<-int,in <-chan int,flag int) {
	c:=0
	//c用来控制将结果放入不同的channel
	//flag标志不同的求平方的goroutine
	for v:=range in{
		u:= v*v
		time.Sleep(500*time.Millisecond)
		//求平方的gorotine1i从对应的channel ch1i中取出数据后
		//再将结果轮流放到ch21,ch22中
		if flag%2==1{
			if c%2==0 {
			out1<- u
		    }else{
			out2<- u
		    }
		}else{
			if c%2==0 {
			out2<- u
		    }else{
			out1<- u
		    }
		}
		c=c+1
	}

	//如果最后一个数据操作完，则关闭ch21,ch22
	if(flag==1){
		close(out1)
		close(out2)
	}
}

func printer(in <-chan int) {
	
	for v:=range in{
		time.Sleep(300*time.Millisecond)
		fmt.Println(v)
		
	}
}

func main() {

   ch11:=make(chan int)
   ch12:=make(chan int)
   ch13:=make(chan int)

   ch21:=make(chan int)
   ch22:=make(chan int)
   t1:=time.Now().UnixNano() 
   
   go counter(ch11,ch12,ch13)

   go squarer(ch21,ch22,ch11,1)
   go squarer(ch21,ch22,ch12,2)
   go squarer(ch21,ch22,ch13,3)

   go printer(ch21)
   printer(ch22)

   t2:=time.Now().UnixNano() 
   fmt.Println("total time is:")
   fmt.Println((t2-t1)/1000000)

}