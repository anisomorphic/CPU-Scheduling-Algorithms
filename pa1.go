// Michael Harris
// mi051467
// COP4600 pa1 - os schedulers

package main

import "os"
import "fmt"
import "bufio"
import "strconv"

type proc struct {
	name string
	arrival int
	burst int
	index int
	selected int
	done int
	preempted int
	waited int
}

//go through processes array and print out any that are arriving this time step (i)
func arrival(array []proc, running []proc, i int, file2 *os.File) ([]proc){
	for j := 0; j < len(array); j++ {
		if array[j].arrival == i {
			fmt.Fprintf(file2, "Time %3d : %s arrived\n", i, array[j].name)
			tempHold := array[j]
			tempHold.arrival = i
			running = append(running, tempHold)
		}
	}
	return running
}

//swap two array positions, used in bubble sorting methods [sort_(.*)]
func swap(array []proc, i, j int) {
	tmp := array[j]
	array[j] = array[i]
	array[i] = tmp
}

//this is very important to the shortest job first mode. in order to accurately
//track wait time, we must not count it if a job ran for an amount of time and
//was preempted. so, we calculate the difference between its original burst
//and current burst, and then += the amount of time it actually ran for into 'preempted'
//if this flag has been set, we decrement by this amount when calculating wait time for a process
//to represent that the process actually did run for x amount of time instead of waiting
func set_preempt(running []proc, currentProc int, processes []proc) {
	tempStr := ""
	tempDiff := 0
	tempUpdate := 0

	for i := 0; i<len(processes); i++ {
		if processes[i].index == currentProc {
			tempStr = processes[i].name
			tempDiff = processes[i].burst
			break
		}
	}

	for i := 0; i<len(running); i++ {
		if running[i].name == tempStr {
			tempUpdate = running[i].burst
			tempDiff -= running[i].burst
			running[i].preempted += tempDiff
			break
		}
	}

//this is important if a process is preempted more than once. we will update
//our processes data structure with the new, updated burst, but we can no longer
//rely on this value for historial purposes (should be fine within scope)
	for i := 0; i<len(processes); i++ {
		if processes[i].index == currentProc {
			processes[i].burst = tempUpdate
			break
		}
	}
}

//bubble sort by arrival time
func sort_arrival(array []proc) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(array) - 1; i++ {
			if array[i + 1].arrival < array[i].arrival {
				swap(array, i, i + 1)
				swapped = true
			}
		}
	}
}

//bubble sort by burst time
func sort_burst(array []proc) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(array) - 1; i++ {
			if array[i + 1].burst < array[i].burst {
				swap(array, i, i + 1)
				swapped = true
			}
		}
	}
}

//bubble sort by index (process number)
func sort_index(array []proc) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(array) - 1; i++ {
			if array[i + 1].index < array[i].index {
				swap(array, i, i + 1)
				swapped = true
			}
		}
	}
}


func main() {

	//make sure we have command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid input file as the first parameter.")
		return
	}
	if len(os.Args) < 3 {
		fmt.Println("Please provide a valid output file as the second parameter.")
		return
	}

	//file management and produce errors if needed
	file2, err2 := os.Create(os.Args[2])
	if err2 != nil {
		fmt.Println("Problem creating output file: ", os.Args[2])
		fmt.Print("\n\t"); panic(err2)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File is invalid: ", os.Args[1])
		fmt.Print("\n\t"); panic(err)
	  return
	}
	defer file.Close()

	//scan input file by word
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	//variables for main
	cmpStr := ""
	processcount := -1
	runfor := -1
	mode := -1
	quantum := -1
	idx := 0
	var processes []proc

	//start handling input file
	for scanner.Scan() {
		inStr := scanner.Text()
		if inStr == "end" {
			break;
		}

		if cmpStr == "processcount" && processcount == -1 {
			processcount, _ = strconv.Atoi(inStr)
		}
		if cmpStr == "runfor" && runfor == -1 {
			runfor, _ = strconv.Atoi(inStr)
		}
		if cmpStr == "use" && mode == -1 {
			if inStr == "rr" {
				mode = 2
			} else if inStr == "sjf" {
				mode = 3
			} else if inStr == "fcfs" {
				mode = 4
			} else {
				panic(inStr)
			}
		}
		if cmpStr == "quantum" && quantum == -1 {
			quantum, _ = strconv.Atoi(inStr)
		}
		if cmpStr == "process" && inStr == "name" {
			scanner.Scan()
			tempName := scanner.Text()

			scanner.Scan(); scanner.Scan(); //"arrival"
			tempArr, _ := strconv.Atoi(scanner.Text())

			scanner.Scan(); scanner.Scan(); //"burst"
			tempBur, _ := strconv.Atoi(scanner.Text())

			tempProc := proc{name:tempName,arrival:tempArr,burst:tempBur,index:idx}
			processes = append(processes, tempProc)
			idx++
		}

		//storing previous word in order to parse
		cmpStr = scanner.Text()
	} //stop processing input

	if (mode == 2) { //RR

			finishedCount := 0
			currentProc := -1
			currentQuantum := -1
			var running, finished []proc

			fmt.Fprintf(file2, "%3d processes\n", processcount)
			fmt.Fprintf(file2, "Using Round-Robin\n")
			fmt.Fprintf(file2, "Quantum %3d\n\n", quantum)

			var time int
			for time = 0; time < runfor; time++ {
				justFinished := false

				//arrived
				running = arrival(processes, running, time, file2)

				//finish
				if currentProc != -1 {
					if len(running) != 0 {
						running[0].burst -= 1
						currentQuantum -= 1

						if running[0].burst == 0 {
							fmt.Fprintf(file2, "Time %3d : %s finished\n", time, running[0].name)
							running[0].done = time
							finished = append(finished, running[0])
							running = running[1:]
							currentProc = -1
							finishedCount++
							justFinished = true
						}
					}
				}

				//select
				if (currentProc == -1) || (currentQuantum < 1){
					if len(running) != 0 {
						if (justFinished == false) {
							tempStruct := running[0]
							running = running[1:]
							running = append(running, tempStruct)
						} else {
							;
						}
						fmt.Fprintf(file2, "Time %3d : %s selected (burst %3d)\n", time, running[0].name, running[0].burst)
						currentProc = running[0].index
						running[0].selected = time
						currentQuantum = quantum
					}
				}

				//account for waiting processes
				for x := 1; x<len(running); x++ {
					running[x].waited++
				}

				if len(running) == 0 {
					fmt.Fprintf(file2, "Time %3d : Idle\n", time)
				}
			}
			fmt.Fprintf(file2, "Finished at time %3d\n\n", time)


			sort_index(finished)
			for i := 0; i < len(finished); i++ {

				fmt.Fprintf(file2, "%s wait %3d turnaround %3d\n", finished[i].name,
												finished[i].waited,
												finished[i].done - finished[i].arrival)
			}
		} //end RR


	if (mode == 3) { //SJF

		sort_burst(processes)
		finishedCount := 0
		currentProc := -1
		var running, finished []proc

		fmt.Fprintf(file2, "%3d processes\n", processcount)
		fmt.Fprintf(file2, "Using preemptive Shortest Job First\n")

		var time int
		for time = 0; time < runfor; time++ {
			//arrived
			running = arrival(processes, running, time, file2)

			//finish
			if currentProc != -1 {
				if len(running) != 0 {
					running[0].burst -= 1

					if running[0].burst == 0 {
						fmt.Fprintf(file2, "Time %3d : %s finished\n", time, running[0].name)
						running[0].done = time
						finished = append(finished, running[0])
						running = running[1:]
						currentProc = -1
						finishedCount++
					}
				}
			}

			//select
			sort_burst(running)
			if (currentProc == -1) || (running[0].index != currentProc) {
				if len(running) != 0 {
					if (running[0].index != currentProc) {
						set_preempt(running, currentProc, processes)
					}
					fmt.Fprintf(file2, "Time %3d : %s selected (burst %3d)\n", time, running[0].name, running[0].burst)
					currentProc = running[0].index
					running[0].selected = time
				}
			}

			if len(running) == 0 {
				fmt.Fprintf(file2, "Time %3d : Idle\n", time)
			}
		}
		fmt.Fprintf(file2, "Finished at time %3d\n\n", time)


		sort_index(finished)
		for i := 0; i < len(finished); i++ {
			ranFor := 0

			if finished[i].preempted != 0 {
				ranFor = finished[i].preempted
			}

			fmt.Fprintf(file2, "%s wait %3d turnaround %3d\n", finished[i].name,
											finished[i].selected - finished[i].arrival - ranFor,
											finished[i].done - finished[i].arrival)
		}
	} //end SJF


	if (mode == 4) { //FCFS

		sort_arrival(processes)
		finishedCount := 0
		currentProc := -1
		var running, finished []proc

		fmt.Fprintf(file2, "%3d processes\n", processcount)
		fmt.Fprintf(file2, "Using First-Come First-Served\n")

		var time int
		for time = 0; time < runfor; time++ {
			//arrived
			running = arrival(processes, running, time, file2)

			//finish
			if currentProc != -1 {
				if len(running) != 0 {
					running[0].burst -= 1

					if running[0].burst == 0 {
						fmt.Fprintf(file2, "Time %3d : %s finished\n", time, running[0].name)
						running[0].done = time
						finished = append(finished, running[0])
						running = running[1:]
						currentProc = -1
						finishedCount++
					}
				}
			}

			//select
			if currentProc == -1 {
				if len(running) != 0 {
					fmt.Fprintf(file2, "Time %3d : %s selected (burst %3d)\n", time, running[0].name, running[0].burst)
					currentProc = running[0].index
					running[0].selected = time
				}
			}

			if len(running) == 0 {
				fmt.Fprintf(file2, "Time %3d : Idle\n", time)
			}
		}
		fmt.Fprintf(file2, "Finished at time %3d\n\n", time)


		sort_index(finished)
		for i := 0; i < len(finished); i++ {
			fmt.Fprintf(file2, "%s wait %3d turnaround %3d\n", finished[i].name,
											finished[i].selected - finished[i].arrival,
											finished[i].done - finished[i].arrival)
		}
	} //end FCFS

	defer file2.Close()
} //end main
