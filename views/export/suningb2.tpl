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

    <title>qulaxin-export</title>
</head>

<body>
    <div id="app">
        <el-container style="height: 500px; border: 1px solid #eee">
            <el-container>
                <el-header style="text-align: right; font-size: 12px">
                    <el-menu :default-active="activeIndex2" class="el-menu-demo" mode="horizontal"
                        @select="handleSelect" background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
                        <el-menu-item index="1"><a href="/qulaxin">功能页</a></el-menu-item>
                    </el-menu>
                </el-header>

                <el-main>
                    <el-card class="box-card">
                        <template #header>
                            <div class="card-header">
                                <span>苏宁B2数据导出</span>
                                <!-- <el-button class="button" type="text">操作按钮</el-button> -->
                            </div>
                        </template>
                        <el-row>
                            <el-col :span="2">
                                <div>时间条件</div>
                            </el-col>
                            <el-col :span="6">
                                <div>
                                    <el-date-picker v-model="timeRange" type="datetimerange" :shortcuts="shortcuts"
                                        language="zh-CN" range-separator="至" start-placeholder="开始日期"
                                        end-placeholder="结束日期" align="right">
                                    </el-date-picker>
                                </div>
                            </el-col>
                            <el-col :span="3">
                                <div>
                                    <el-button type="primary" @click.stop="exportTap" :loading="exportBtnLoading">导出
                                    </el-button>
                                </div>
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

                    exportBtnLoading: false,
                    shortcuts: [{
                        text: '今天',
                        value: (() => {
                            const end = new Date();
                            const start = new Date();
                            start.setHours(0, 0, 0, 0)
                            end.setTime(start.getTime() + 3600 * 1000 * 24 * 1 - 1);
                            return [start, end]
                        })()
                    }, {
                        text: '昨天',
                        value: (() => {
                            const end = new Date();
                            const start = new Date();
                            start.setHours(0, 0, 0, 0)
                            start.setTime(start.getTime() - 3600 * 1000 * 24 * 1);
                            end.setTime(start.getTime() + 3600 * 1000 * 24 * 1 - 1);
                            return [start, end]
                        })()
                    }, {
                        text: '前天',
                        value: (() => {
                            const end = new Date();
                            const start = new Date();
                            start.setHours(0, 0, 0, 0)
                            start.setTime(start.getTime() - 3600 * 1000 * 24 * 2);
                            end.setTime(start.getTime() + 3600 * 1000 * 24 * 1 - 1);
                            return [start, end]
                        })()
                    }, {
                        text: '三天前',
                        value: (() => {
                            const end = new Date();
                            const start = new Date();
                            start.setHours(0, 0, 0, 0)
                            start.setTime(start.getTime() - 3600 * 1000 * 24 * 3);
                            end.setTime(start.getTime() + 3600 * 1000 * 24 * 1 - 1);
                            return [start, end]
                        })()
                    }],
                    timeRange: ''
                };
            },
            methods: {
                exportTap() {
                    if (!this.timeRange) {
                        this.$message.warning({
                            message: '请选择时间条件',
                            type: 'warning'
                        });
                        return
                    }

                    if (this.timeRange[1].getTime() - this.timeRange[0].getTime() > 3600 * 1000 * 24 * 1) {
                        this.$message.warning({
                            message: '时间范围不能大于一天',
                            type: 'warning'
                        });
                        return
                    }

                    this.exportBtnLoading = true
                    location.href = "/export/export-suning-b2?start_time=" + dateFormat("YYYY-mm-dd HH:MM:SS", this.timeRange[0]) + "&end_time=" + dateFormat("YYYY-mm-dd HH:MM:SS", this.timeRange[1])
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
        .card-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
        }

        .box-card {
            width: 100%;
        }
    </style>
</body>

</html>