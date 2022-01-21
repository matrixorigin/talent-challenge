#### Description

This is a simple example that performs an operation equivalent to the following sql:

```sql
select (a - 99.88 + b + c) as A from R order by A;
```

You need to modify the example (or even refactor the overall code base) to speed up the entire execution process, with the caveat that you cannot rely on the order of the input data. The modified program needs to be loadable via a csv file, with the csv corresponding to the following data example:

```csv
1, 2, 3.1
2, 4, 5.1
```

The above csv corresponds to the following schema:

```sql
create table R(a int, b int, c float);
```

The results need to be exported to csv as well, ensuring a performance improvement of several times to above an order of magnitude, and indicating why the original execution process was slow and how it was accelerated. 

Please submit your solution to  Yingfeng DOT Zhang AT matrixorigin DOT cn via email

**WARNING: DO NOT DIRECTLY FORK THIS REPO. DO NOT PUSH PROJECT SOLUTIONS PUBLICLY.**
