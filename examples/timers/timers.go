// Часто бывает нужно выполнить код Go в какой-то момент
// в будущем или выполнить код несколько раз с некоторым
// интервалом. Встроенные в Go _timer_ и _ticker_ функции
// позволяют легко решить обе задачи.
// Сначала посмотрим на таймер, затем на
// [счетчик тиков](tickers).

package main

import "time"
import "fmt"

func main() {

    // Таймер представляет из себя одиночное событие в будущем.
    // Вы говорите таймеру как долго нужно подождать и он
    // обеспечивает канал, который будет оповещён в указанное
    // время. Этот таймер будет ждать 2 секунды.
    timer1 := time.NewTimer(time.Second * 2)

    // `<-timer1.C` блокирует канал таймера `C`
    // до тех пор, пока он не отправит значение, означающее,
    // что вышел срок таймера.
    <-timer1.C
    fmt.Println("Timer 1 expired")

    // Если нужно просто подождать, можно использовать
    // `time.Sleep`. Причина, по которой может быть полезен
    // таймер в том, что можно отменить таймер до окончания
    // его срока. Вот пример отмены.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
