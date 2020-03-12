rm ex
rm *.stu
go build -o ex $1
# do 2 processes
echo "c2-rr"
./ex c2-rr.in c2-rr.stu
diff c2-rr.stu c2-rr.base
# now do five processes
echo "c5-rr"
./ex c5-rr.in c5-rr.stu
diff c5-rr.stu c5-rr.base
# now do 10 processes
echo "c10-rr"
./ex c10-rr.in c10-rr.stu
diff c10-rr.stu c10-rr.base
echo "c2-sjf"
./ex c2-sjf.in c2-sjf.stu
diff c2-sjf.stu c2-sjf.base
echo "c5-sjf"
./ex c5-sjf.in c5-sjf.stu
diff c5-sjf.stu c5-sjf.base
echo "c10-sjf"
./ex c10-sjf.in c10-sjf.stu
diff c10-sjf.stu c10-sjf.base
echo "c2-fcfs"
./ex c2-fcfs.in c2-fcfs.stu
diff c2-fcfs.stu c2-fcfs.base
echo "c5-fcfs"
./ex c5-fcfs.in c5-fcfs.stu
diff c5-fcfs.stu c5-fcfs.base
echo "c10-fcfs"
./ex c10-fcfs.in c10-fcfs.stu
diff c10-fcfs.stu c10-fcfs.base
