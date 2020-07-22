num=int(input("enter the number"))
temp=num
su=0
while (temp>0):
    digit=temp%10
    su=su+(digit**3)
    temp=temp//10
print(su)
if su==num:
    print("you have entered an amstrong number")
else:
    print("not an amstrong number")