echo "c2-sjf"
./ex c2-sjf.in c2-sjf.stu
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
diff c2-sjf.stu c2-sjf.base
echo "c5-sjf"
./ex c5-sjf.in c5-sjf.stu
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
diff c5-sjf.stu c5-sjf.base
echo "c10-sjf"
./ex c10-sjf.in c10-sjf.stu
rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
diff c10-sjf.stu c10-sjf.base
