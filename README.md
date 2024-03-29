# F1 Race Analysis

This application analysis the results of a race by a log file containing each driver lap status.  The application is a command-line application built-in Golang and all the data is persisted on the memory of the runtime. 

## Input data

When the application statrts is going to be promped to inform a few information needed to
perform the race log analysis:

1. <b>Path of the file</b>: Here you must inform where on your OS is stored the file that this application must analyze. 

The file must be with the following format:

```text
Local Time        Racer                                   Lap             Lap Time          Average Speed on the lap
23:49:08.277      038 – F.MASSA                           1		1:02 852                        44,275
23:49:10.858      033 – R.BARRICHELLO                     1		1:04.352                        43,243
23:49:11.075      002 – K.RAIKKONEN                       1             1:04.108                        43,408
23:49:12.667      023 – M.WEBBER                          1		1:04.414                        43,202
23:49:30.976      015 – F.ALONSO                          1		1:18.456			35,47
23:50:11.447      038 – F.MASSA                           2		1:03.170                        44,053
23:50:14.860      033 – R.BARRICHELLO                     2		1:04.002                        43,48
23:50:15.057      002 – K.RAIKKONEN                       2             1:03.982                        43,493
23:50:17.472      023 – M.WEBBER                          2		1:04.805                        42,941
23:50:37.987      015 – F.ALONSO                          2		1:07.011			41,528
23:51:14.216      038 – F.MASSA                           3		1:02.769                        44,334
23:51:18.576      033 – R.BARRICHELLO		          3		1:03.716                        43,675
23:51:19.044      002 – K.RAIKKONEN                       3		1:03.987                        43,49
23:51:21.759      023 – M.WEBBER                          3		1:04.287                        43,287
23:51:46.691      015 – F.ALONSO                          3		1:08.704			40,504
23:52:01.796      011 – S.VETTEL                          1		3:31.315			13,169
23:52:17.003      038 – F.MASS                            4		1:02.787                        44,321
23:52:22.586      033 – R.BARRICHELLO		          4		1:04.010                        43,474
23:52:22.120      002 – K.RAIKKONEN                       4		1:03.076                        44,118
23:52:25.975      023 – M.WEBBER                          4		1:04.216                        43,335
23:53:06.741      015 – F.ALONSO                          4		1:20.050			34,763
23:53:39.660      011 – S.VETTEL                          2		1:37.864			28,435
23:54:57.757      011 – S.VETTEL                          3		1:18.097			35,633

```

>Note:
>In the file you must remove the header line.

2. <b>Total number of laps on the race</b>: This is used to know if a racer ended or not the race. 

![image](https://user-images.githubusercontent.com/40045069/211218380-25dec1c4-14fe-4bc4-81a9-f1c14218118b.png)

## Output

After the execution, is going to be printing something like this on the console:

![image](https://user-images.githubusercontent.com/40045069/220788885-dd1aef73-1277-4071-b608-7156a0fa427f.png)

## Configuration

To run this application you can run it by Docker or with golang (you must have it installed on your machine).

The commands to run by golang are:

```bash
go mod download
go install
go build

./f1-race-analysis
```

Or by Docker:

```bash
docker build -t f1-race-analysis:latest .
docker run -it f1-race-analysis:latest
```
