package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Muhammad Dimas Afifudin" //mengirim data ke channel
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel // mengambil data dari channel dan dimasukkan ke variabel data
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Dimas Afifudin"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel // mengambil data dari channel dan dimasukkan ke variabel data
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Muhammad Dimas Afifudin"
} //function yang hanya bisa mengirim data

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
} // hanya bisa mengambil data dari channel

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	func() {
		channel <- "Muhammad"
		channel <- "Dimas"
		channel <- "Afifudin"

	}()

	func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Dari Channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}
}
