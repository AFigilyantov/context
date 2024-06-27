/*
1. Напишите 2 функции: 1 функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter.
Вторая функция пишет эти данные в файл. Передайте в эти функции контекст.
Используйте context и waitgroup.
*/

package main

func main() {

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// go mainWork(ctx)

	// <-ctx.Done()
	// fmt.Println("ctx.Done")
	// os.Exit(0)

}

// func mainWork(ctx context.Context) {
// 	wg := &sync.WaitGroup{}
// 	readConsChan := make(chan string)

// 	wg.Add(2)
// 	go readConsole(ctx, wg, readConsChan)
// 	go writeToFile(ctx, wg, readConsChan)
// 	wg.Wait()

// }

// func readConsole(ctx context.Context, wg *sync.WaitGroup, ch chan string) <-chan string {
// 	sc := bufio.NewScanner(os.Stdin)
// 	go func() {

// 		fmt.Println("Пиши уже. 5 секунд. Время пошло...")
// 		defer close(ch)
// 		for sc.Scan() {
// 			txt := sc.Text()
// 			ch <- txt

// 		}
// 		defer wg.Done()

// 	}()

// 	return ch
// }

// func writeToFile(ctx context.Context, wg *sync.WaitGroup, ch <-chan string) {

// 	for txt := range ch {
// 		writeText(ctx, txt)

// 	}
// 	defer wg.Done()
// }

// func writeText(ctx context.Context, text string) {
// 	file, err := os.OpenFile("savedTextFromKeyboard.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

// 	if err != nil {
// 		fmt.Println("Unable to create  or append to file:", err)
// 		os.Exit(1)
// 	}

// 	file.WriteString(fmt.Sprintf("%s\n", text))

// 	fmt.Println("Writen.")
// 	defer file.Close()
// }
