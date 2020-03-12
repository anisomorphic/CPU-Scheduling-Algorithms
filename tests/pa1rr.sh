go build -o ex $1
# do 2 processes
echo "c2-rr"
./ex c2-rr.in c2-rr.stu
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
diff c2-rr.stu c2-rr.base
# now do five processes
echo "c5-rr"
./ex c5-rr.in c5-rr.stu
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
diff c5-rr.stu c5-rr.base
# now do 10 processes
echo "c10-rr"
./ex c10-rr.in c10-rr.stu
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
diff c10-rr.stu c10-rr.base
