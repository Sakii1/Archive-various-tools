import subprocess
import tkinter as tk
import chardet
import os
from tkinter import messagebox



class CustomApp(tk.Tk):

    def __init__(self, input1_width, input1_height, input2_width, input2_height,
                 input4_width, input4_height, title):
        super().__init__()  # 继承 tk.Tk 类
        self.title(title)  # 设置窗口标题

        # 创建文本变量，用于存储输入框的值
        self.input1_value = tk.StringVar()
        self.input2_value = tk.StringVar()
        self.input4_value = tk.StringVar()

        self.input2_value.set(130)

        # 创建标签和输入框，并设置长宽以及 font
        self.input1_label = tk.Label(self, text="(可配合new使用)装扮ID：")

        self.input1_label.pack()
        self.input1_entry = tk.Entry(self, width=input1_width, font=('Arial', input1_height),
                                     textvariable=self.input1_value)
        self.input1_entry.pack()

        self.input1_label.place(x=200, y=250)
        self.input1_entry.place(x=350, y=250)

        self.input2_label = tk.Label(self, text="录制频率(ms)：", font=("Arial", 12))
        self.input2_label.pack()
        self.input2_entry = tk.Entry(self, width=input2_width, font=('Arial', input2_height),
                                     textvariable=self.input2_value)
        self.input2_entry.pack()

        self.input2_label.place(x=0, y=250)
        self.input2_entry.place(x=120, y=250)

        self.input4_label = tk.Label(self, text="在下方放入cookie的值 (value),第一次需要保存.之后使用可以留空 直至过期", font=("Arial", 11))  # 12 = 字体大小
        self.input4_label.pack()
        self.input4_entry = tk.Entry(self, width=input4_width, font=('Arial', input4_height),
                                     textvariable=self.input4_value)
        self.input4_label.place(x=0, y=300)
        self.input4_entry.place(x=50, y=330)

        # 创建一个按钮
        self.button = tk.Button(self, text="开始录制", fg="red", command=self.on_click, width=20, height=3)
        self.button.pack()
        self.button.place(x=10, y=10)

        self.query_button = tk.Button(self, text="暂停录制", fg="blue", command=self.on_query, width=15, height=2)
        self.query_button.place(x=380, y=10)  # 将按钮放置在(50, 20)坐标

        self.query_button2 = tk.Button(self, text="继续录制", fg="green", command=self.on_query2, width=15, height=2)
        self.query_button2.place(x=380, y=120)  # 将按钮放置在(50, 20)坐标

        self.query_button3 = tk.Button(self, text="终止录制，开始生成排序", fg="#FF00FF", command=self.on_query3, width=20,
                                       height=3)
        self.query_button3.place(x=10, y=150)  # 将按钮放置在(50, 20)坐标

        self.query_button4 = tk.Button(self, text="保存cookie", fg="#6A5ACD", command=self.on_query4, width=15, height=2)
        self.query_button4.place(x=380, y=330)  # 将按钮放置在(50, 20)坐标




    def on_query(self):  # 暂停录制
        with open('config.txt', 'rb') as f:
            result = chardet.detect(f.read())
            encoding = result['encoding']

        with open("config.txt", "r+", encoding=encoding) as file:
            lines = file.readlines()  # 读取文件中的所有行

            lines[4 - 1] = "2" + "\n"  # 在第三行写入

            file.seek(0)  # 将文件指针移回文件开头
            file.writelines(lines)  # 将修改后的内容写回文件

        pass

    def on_query2(self):  # 继续录制
        with open('config.txt', 'rb') as f:
            result = chardet.detect(f.read())
            encoding = result['encoding']

        with open("config.txt", "r+", encoding=encoding) as file:
            lines = file.readlines()  # 读取文件中的所有行

            lines[4 - 1] = "1" + "\n"  # 在第三行写入

            file.seek(0)  # 将文件指针移回文件开头
            file.writelines(lines)  # 将修改后的内容写回文件


    def on_query4(self):  # 保存cookie
        cookiess = self.input4_value.get()


        with open('config.txt', 'rb') as f:
            result = chardet.detect(f.read())
            encoding = result['encoding']

        with open("config.txt", "r+", encoding=encoding) as file:
            lines = file.readlines()  # 读取文件中的所有行

            lines[3 - 1] = cookiess + "\n"  # 在第二行写入

            file.seek(0)  # 将文件指针移回文件开头
            file.writelines(lines)  # 将修改后的内容写回文件

            messagebox.showinfo("", "cookie保存成功\n")


    def on_query3(self):  # 终止录制,排序+生成
        with open('config.txt', 'rb') as f:
            result = chardet.detect(f.read())
            encoding = result['encoding']

        with open("config.txt", "r+", encoding=encoding) as file:
            lines = file.readlines()  # 读取文件中的所有行

            lines[4 - 1] = "3" + "\n"  # 在第三行写入

            file.seek(0)  # 将文件指针移回文件开头
            file.writelines(lines)  # 将修改后的内容写回文件

    def on_click(self):  # 开始录制

        item_id = self.input1_value.get()
        ms = self.input2_value.get()

        # print(item_id,ms,cookiess)

        # line_number = 3  # 要写入或覆盖信息的行数

        with open('config.txt', 'rb') as f:
            result = chardet.detect(f.read())
            encoding = result['encoding']

        with open("config.txt", "r+", encoding=encoding) as file:
            lines = file.readlines()  # 读取文件中的所有行
            lines[2 - 1] = item_id + "\n"  # 在第一行写入


            lines[4 - 1] = "1" + "\n"  # 在第三行写入

            lines[5 - 1] = ms + "\n"  # 在第四行写入
            file.seek(0)  # 将文件指针移回文件开头
            file.writelines(lines)  # 将修改后的内容写回文件

        # 获取当前目录
        current_dir = os.getcwd()

        # exe程序的路径
        exe_path = os.path.join(current_dir, 'luzhi.exe')

        # 构建cmd命令
        cmd_command = f'start {exe_path}'

        # 调用cmd命令行打开exe程序
        subprocess.call(cmd_command, shell=True)


# 创建应用实例
if __name__ == "__main__":
    app = CustomApp(10, 17, 5, 15, 20, 20, "录号图形控制界面")
    app.geometry("530x400")
    # app.on_query()
    app.mainloop()
