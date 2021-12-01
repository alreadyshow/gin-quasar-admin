import { getAction, postAction, putAction } from 'src/api/manage'
import { FormatDataTime } from 'src/utils/date'
import { mapActions } from 'vuex'

export const addOrEditMixin = {
    computed: {
        formTypeName() {
            if (this.formType === 'edit') {
                return '编辑'
            } else if (this.formType === 'add') {
                return '新增'
            } else {
                return '错误'
            }
        },
        showDateTime() {
            return (datetime) => {
                return FormatDataTime(datetime)
            }
        }
    },
    data() {
        return {
            addOrEditDetail: {},
            addOrEditVisible: false,
            formType: '',
            loading: false,
            options: {},
        }
    },
    async created() {
        const detailLocal = this.$q.localStorage.getItem('gqa-dict')
        if (detailLocal) {
            this.options = detailLocal
        } else {
            await this.GetGqaDict()
            this.options = this.$q.localStorage.getItem('gqa-dict')
        }
    },
    methods: {
        ...mapActions('storage', ['GetGqaDict']),
        show(row) {
            this.loading = true
            this.resetDetail()
            this.addOrEditDetail = Object.assign(this.detail, row)
            this.addOrEditVisible = true
            if (!this.addOrEditDetail.id) {
                this.loading = false
            } else {
                this.handleQueryById(this.addOrEditDetail.id)
            }
        },
        async handleQueryById(id) {
            const res = await postAction(this.url.queryById, {
                id,
            })
            if (res.code === 1) {
                this.addOrEditDetail = res.data.records
            }
            this.loading = false
        },
        async handleAddOrEidt() {
            const success = await this.$refs.addOrEditForm.validate()
            if (success) {
                // if (this.addOrEditDetail.avatar) {
                //     this.addOrEditDetail.avatar = JSON.stringify(this.addOrEditDetail.avatar)
                // }
                if (this.formType === 'edit') {
                    if (this.url === undefined || !this.url.edit) {
                        this.$q.notify({
                            type: 'negative',
                            message: "请先配置url",
                        })
                        return
                    }
                    const res = await putAction(this.url.edit, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.addOrEditVisible = false
                    }
                } else if (this.formType === 'add') {
                    if (this.url === undefined || !this.url.add) {
                        this.$q.notify({
                            type: 'negative',
                            message: "请先配置url",
                        })
                        return
                    }
                    const res = await postAction(this.url.add, this.addOrEditDetail)
                    if (res.code === 1) {
                        this.$q.notify({
                            type: 'positive',
                            message: res.message,
                        })
                        this.addOrEditVisible = false
                    }
                } else {
                    this.$q.notify({
                        type: 'negative',
                        message: '无法新增或编辑！',
                    })
                }
                this.$emit('handleFinish')
            } else {
                this.$q.notify({
                    type: 'negative',
                    message: '请完善表格信息！',
                })
            }
        },
    }
}