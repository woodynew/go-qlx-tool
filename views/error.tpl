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

    <title>error</title>
</head>

<body>
    <div id="app">
        <el-container style="height: 500px; border: 1px solid #eee">
            <el-container>
                <el-main>
                    <el-card class="box-card" :body-style="{height:'80%'}">
                        <el-row type="flex" justify="center">
                            <el-col :span="12">
                                <el-alert title="{{b .Message b}}" type="warning" show-icon center :closable="false"
                                    style="height: 60px;">
                                </el-alert>

                            </el-col>
                        </el-row>
                        <el-row type="flex" justify="center">
                            <el-col :span="12" style="text-align: center;margin-top: 20px;">
                                <el-link type="info" underline href="{{b .RetUrl b}}">点击返回
                                </el-link>
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
                };
            },
            methods: {

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