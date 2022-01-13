<html>

<head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1.0" />

    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/element-plus/dist/index.css" />
    <!-- Import Vue 3 -->
    <script src="https://cdn.jsdelivr.net/npm/vue@3"></script>
    <!-- Import component library -->
    <script src="https://cdn.jsdelivr.net/npm/element-plus"></script>
    <script src="https://cdn.jsdelivr.net/npm/element-plus/lib/locale/lang/zh-cn.js"></script>

    <script src="https://cdn.jsdelivr.net/npm/jquery@3/dist/jquery.min.js"></script>

    <!-- <script src="https://unpkg.com/dayjs/locale/zh-cn.js"></script> -->

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

        const app = Vue.createApp(App);
        app.use(ElementPlus, {
            locale: zhCn,
        });
        app.mount("#app");

    </script>

    <style>
        .box-card {
            width: 100%;
        }
    </style>
</body>

</html>