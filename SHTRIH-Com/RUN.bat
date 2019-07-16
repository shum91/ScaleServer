:: This example run script

shtrih-com.exe ":50505:COM1:9600" 1>> "%Date:~6,4%%Date:~3,2%%Date:~0,2%_out.log" 2>>&1

:: not log
:: shtrih-com.exe ":50505:COM1:9600" 1>> NUL 2>>&1

pause