import tkinter as tk
from tkinter import messagebox
import requests
import urllib3;

urllib3.disable_warnings()
# import chardet

# with open('data.txt', 'rb') as f:
#     result = chardet.detect(f.read())
#     encoding = result['encoding']
#
# with open('data.txt', 'r', encoding=encoding) as file:
#     lines = file.readlines()
#     item_id = lines[1]
#     cooki = lines[3]


class CustomApp(tk.Tk):

    def __init__(self, input1_width, input1_height, input2_width, input2_height, input3_width, input3_height,
                 input4_width, input4_height, title):
        super().__init__()  # 继承 tk.Tk 类
        self.title(title)  # 设置窗口标题

        # 创建文本变量，用于存储输入框的值
        self.input1_value = tk.StringVar()
        self.input2_value = tk.StringVar()
        self.input3_value = tk.StringVar()
        self.input4_value = tk.StringVar()

        # self.input1_value.set('{}'.format(item_id))
        # self.input4_value.set('{}'.format(cooki))

        # 创建标签和输入框，并设置长宽以及 font
        self.input1_label = tk.Label(self, text="item_id：")
        self.input1_label.pack()
        self.input1_entry = tk.Entry(self, width=input1_width, font=('Arial', input1_height),
                                     textvariable=self.input1_value)
        self.input1_entry.pack()

        self.input2_label = tk.Label(self, text="对方uid：")
        self.input2_label.pack()
        self.input2_entry = tk.Entry(self, width=input2_width, font=('Arial', input2_height),
                                     textvariable=self.input2_value)
        self.input2_entry.pack()

        self.input3_label = tk.Label(self, text="赠送的编号：")
        self.input3_label.pack()
        self.input3_entry = tk.Entry(self, width=input3_width, font=('Arial', input3_height),
                                     textvariable=self.input3_value)
        self.input3_entry.pack()

        self.input4_label = tk.Label(self, text="cookie：")
        self.input4_label.pack()
        self.input4_entry = tk.Entry(self, width=input4_width, font=('Arial', input4_height),
                                     textvariable=self.input4_value)
        self.input4_entry.pack()

        self.input1_label5 = tk.Label(self, text="可以先填uid查询 ↓")
        self.input1_label5.place(x=300, y=50)

        # 创建一个按钮
        self.button = tk.Button(self, text="发送装扮", command=self.on_click, width=20, height=3)
        self.button.pack()

        self.query_button = tk.Button(self, text="查询该UID的昵称", command=self.on_query)
        self.query_button.place(x=300, y=80)  # 将按钮放置在(50, 20)坐标

        self.query_button2 = tk.Button(self, text="更换佩戴编号", command=self.on_query2)
        self.query_button2.place(x=300, y=120)  # 将按钮放置在(50, 20)坐标

        self.query_button3 = tk.Button(self, text="查询剩余编号", command=self.on_query3)
        self.query_button3.place(x=50, y=120)  # 将按钮放置在(50, 20)坐标



    def on_query(self):

        uid = self.input2_value.get()
        url = f"https://account.bilibili.com/api/member/getCardByMid?mid={uid}"
        response = requests.get(url)
        data = response.json()
        name = data['card']['name']
        messagebox.showinfo("查询结果", f"用户名:{name}\nUID:{uid}\n")

    def on_query2(self):

        item_id = self.input1_value.get()

        number = self.input3_value.get()
        cooki = self.input4_value.get()
        #
        cookie = {}
        for item in cooki.split(";"):
            if "=" in item:
                key, value = item.split("=")
                cookie[key.strip()] = value.strip()

        csrf = cookie.get('bili_jct')
        url = "https://api.bilibili.com/x/garb/user/fannum/change"

        data = {
            "item_id": "{}".format(int(item_id)),
            "num": "{}".format(number),
            "csrf": "{}".format(csrf),
        }

        headers = {
            "native_api_from": "h5",
            "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
            # "Cookie":f"{ck}",
            "Host": "api.bilibili.com",
            "Connection": "Keep-Alive",
            "Accept-Encoding": "gzip",
            "Content-Length": "100",
        }
        response = requests.post(url, cookies=cookie, headers=headers, data=data)
        data = response.json()
        name = data['message']
        messagebox.showinfo("更换结果", f"{name}\n")

    def on_query3(self):
        item_id = self.input1_value.get()

        cooki = self.input4_value.get()
        #
        cookie = {}
        for item in cooki.split(";"):
            if "=" in item:
                key, value = item.split("=")
                cookie[key.strip()] = value.strip()

        url = "https://api.bilibili.com/x/garb/user/fannum/list?item_id={}".format(item_id)

        headers = {
            "native_api_from": "h5",
            "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
            # "Cookie":f"{ck}",
            "Host": "api.bilibili.com",
            "Connection": "Keep-Alive",

        }
        response = requests.get(url, cookies=cookie, headers=headers)
        data = response.json()
        dataa = response.text
        print(dataa)

        jishu = dataa.count("buy_mid")
        if jishu != 0:
            bh = []
            for i in range(0, jishu):
                numberr = data['data']['list'][i]['number']
                bh.append(numberr)

            print("编号:", bh)
            messagebox.showinfo("编号", f"{bh}\n")
        else:
            print(data)


    def on_click(self):
        # 从 4 个输入框中获取输入值
        # global csrf
        item_id = self.input1_value.get()
        uid = self.input2_value.get()
        number = self.input3_value.get()
        cooki = self.input4_value.get()
        #
        cookie = {}
        for item in cooki.split(";"):
            if "=" in item:
                key, value = item.split("=")
                cookie[key.strip()] = value.strip()
        print(item_id)
        print(cookie)
        csrf = cookie.get('bili_jct')
        print(csrf)

        url = "https://api.bilibili.com/x/garb/user/fannum/present/v2"

        data = {
            "item_id": "{}".format(int(item_id)),
            "fan_num": "{}".format(number),
            "to_mid": "{}".format(uid),
            "csrf": "{}".format(csrf),
        }

        # print(data)

        headers = {
            "native_api_from": "h5",
            # "Cookie": f"{cooki}",
            "Referer": "https://www.bilibili.com/h5/mall/send/home/{}?navhide=1".format(int(item_id)),
            "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
            "Content-Length": "{}".format(len(data)),
            "X-CSRF-TOKEN": "{}".format(csrf),
            "Host": "api.bilibili.com",
            "Connection": "Keep-Alive",
            "Accept-Encoding": "gzip"
        }
        # print(url,data,headers)
        response = requests.post(url, cookies=cookie, headers=headers, data=data)
        if response.json()['code'] == 0:
            messagebox.showinfo("结果","发送成功")
        else:
            messagebox.showinfo("结果",
                                f"发送失败：{response.text}")

        print(response.text)

        # messagebox.showinfo("end:",
        #                     f"输入框1的内容是：{response.json()}")

        # 弹出消息框，输出 4 个输入框中的值


# 创建应用实例
if __name__ == "__main__":
    app = CustomApp(10, 20, 10, 20, 10, 20, 20, 30, "编号发送小工具2.0")
    # app.on_query()
    app.mainloop()
