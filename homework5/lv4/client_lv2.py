import requests
import json
students={
        'student1':{
        'name':'Tom',
        'id':'2024214000',
        'address':'China',
        'birth':'',
        'gender':'Male',}
        ,'student2':{
        'name':'Jack',
        'id':'2024214001',
        'address':'US',
        'birth':'2014-05-14',
        'gender':'Male',}
            }
for key,value in students.items():
    try:
        r=requests.post("http://127.0.0.1:8888/add",data=json.dumps(value))
    except Exception as e:
        print(key,e)
    else:
        data=r.json()
        print(key,data)
        print()
        print()
student1={
        'name':'Tom',
        'id':'2024214000',
        'address':'China',
        'birth':'2010-06-18',
        'gender':'Male',
        }
try:
    r=requests.post("http://127.0.0.1:8888/profile",data=json.dumps(student1))
except Exception as e:
    print(e)
else:
        data=r.json()  
        print(key,data)
        print()
        print()
studentid="2024214000"
try:
    r=requests.get("http://127.0.0.1:8888/search?id="+studentid)
except Exception as e :
    print(e)
else:
        data=r.json()
        print(key,data)
        print()
        print()




