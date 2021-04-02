<html>

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1.0" />
    <script src="https://unpkg.com/vue@next"></script>
    <!-- import CSS -->
    <link rel="stylesheet" href="https://unpkg.com/element-plus/lib/theme-chalk/index.css">
    <!-- import JavaScript -->
    <script src="https://unpkg.com/element-plus/lib/index.full.js"></script>
    <script src="https://unpkg.com/element-plus/lib/umd/locale/zh-cn.js"></script>
    <script src="https://unpkg.com/dayjs/locale/zh-cn.js"></script>

    <script src="/static/js/utils.js"></script>

    <title>qulaxin</title>
</head>

<body>
    <div id="app">
        <el-container style="height: 500px; border: 1px solid #eee">
            <el-container>

                <el-main>
                    <el-card class="box-card">
                        <el-row type="flex" justify="center">
                            <el-col :span="12">
                                <el-button style="width: 100%;" type="primary" plain @click.stop="btn1Tap">
                                    苏宁B2数据导出
                                </el-button>
                            </el-col>
                        </el-row>
                    </el-card>
                </el-main>
            </el-container>
        </el-container>
    </div>
    <script>

        const App = {
            data() {
                return {
                    message: "Hello Element Plus",


                };
            },
            methods: {
                btn1Tap() {
                    location.href = "/export/get-suning-b2"
                }
            },
        };
        ElementPlus.locale(ElementPlus.lang.zhCn)

        const app = Vue.createApp(App);
        app.use(ElementPlus);
        app.mount("#app");

    </script>

    <style>
        .box-card {
            width: 100%;
        }
    </style>
</body>

</html>