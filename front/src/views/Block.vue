<template>
  <div style="display: flex;justify-content: center;">
    <!-- <el-main> -->
    <div style="width:900px">
      <el-main>
        <el-upload
          class=""
          drag
          :action="action"
          name="block"
          :on-success="success"
          :on-error="onError"
        >
          <i class="el-icon-upload"></i>
          <div class="el-upload__text">
            将区块文件拖到此处，或<em>点击上传</em>
          </div>
          <div class="el-upload__tip" slot="tip">
            上传通过 'peer channel fetch'
            命令或其他方式从网络中拉取的单个区块文件(eg：mychannel.block)
          </div>
        </el-upload>
        <el-divider content-position="center">OR</el-divider>
        <el-form ref="form" :model="sizeForm">
          <el-form-item label="" prop="">
            <el-input
              type="textarea"
              :rows="10"
              v-model="blockForm.block"
              placeholder="输入单个区块数据的 base64 编码字符串"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submit">提交</el-button>
          </el-form-item>
        </el-form>
      </el-main>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      action: `${this.$consts.ServerAddrPre}/block/file`,
      blockForm: {
        block: ""
      }
    };
  },
  methods: {
    success(response, file, fileList) {
      console.log(response);
      if (response.code == 200) {
        this.handleBlock(response.data);
      } else {
        this.$message({
          message: "区块解析错误，请检查上传文件是否正确",
          type: "warnning"
        });
      }
    },
    onError(err, file, fileList) {
      this.$message.error("上传失败");
    },
    submit() {
      const post = `${this.$consts.ServerAddrPre}/block/raw`;
      const formdata = new FormData();
      formdata.append("block", this.blockForm.block);
      this.$axios
        .post(post, formdata)
        .then(rsp => {
          if (rsp.data.code === 200) {
            this.handleBlock(rsp.data.data);
          } else {
            this.$message({
              message: "请求数据出错：" + rsp.data.msg,
              type: "warnning"
            });
          }
        })
        .catch(err => {
          this.$message.error(err.message);
        });
    },
    handleBlock(block) {
      localStorage.block = JSON.stringify(block);
      if (block.type == 1) {
        this.$router.push({
          name: "Config"
        });
      } else if (block.type == 0) {
        this.$router.push({
          name: "Transaction"
        });
      }
    }
  }
};
</script>

<style>
</style>