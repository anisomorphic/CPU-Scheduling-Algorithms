echo "c2-fcfs"
./ex c2-fcfs.in c2-fcfs.stu
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
diff c2-fcfs.stu c2-fcfs.base
echo "c5-fcfs"
./ex c5-fcfs.in c5-fcfs.stu
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
diff c5-fcfs.stu c5-fcfs.base
echo "c10-fcfs"
./ex c10-fcfs.in c10-fcfs.stu
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
diff c10-fcfs.stu c10-fcfs.base
