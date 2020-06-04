@echo on
::bee generate scaffold goods -fields="id:int,name:string,image:string" -conn="root:123456@tcp(127.0.0.1:3306)/we7"
::bee generate appcode -tables="ims_shimmer_liveshop_lives" -conn="root:123456@tcp(127.0.0.1:3306)/we7"


echo text input

set project=
set /p project=:
set input=
set /p input=:
echo %input% is input %project% is project
cd %input%
rem @echo on
for  %%a in (*) do (
    echo %%~na
	cd %project%
	del /f d:\go\src\live\routers\router.go
	bee generate appcode -tables="%%~na" -conn="root:123456@tcp(127.0.0.1:3306)/we7"  
)


pause