<template>
  <el-card style="margin-top:5px;">
    <el-row style="border-bottom:dotted;border-width:1px;" class="border">
      <div style="margin-bottom:5px">
        <el-row>
          <el-col :span="4" class="label">提案请求MSPID：</el-col>
          <el-col :span="20" class="text">{{
            transaction.signer.mspid
          }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="4" class="label">提案请求者证书：</el-col>
          <el-col :span="20" class="text">
            <el-popover placement="right-start" trigger="hover">
              <span class="code_text">{{ transaction.signer.cert }}</span>
              <i class="el-icon-bank-card" slot="reference"></i>
            </el-popover>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="4" class="label">提案请求者签名：</el-col>
          <el-col :span="20" class="text">{{ transaction.signature }}</el-col>
        </el-row>
      </div>
      <div>
        <el-row>
          <el-col :span="4" class="label">通道：</el-col>
          <el-col :span="20" class="text">{{ transaction.channel }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="4" class="label">TxID：</el-col>
          <el-col :span="20" class="text">{{ transaction.tx_id }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="4" class="label">时间：</el-col>
          <el-col :span="20" class="text">{{ transaction.time }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="4" class="label">链码：</el-col>
          <el-col :span="20" class="text">{{ transaction.chaincode }}</el-col>
        </el-row>
        <el-row>
          <el-col :span="4" class="label">方法：</el-col>
          <el-col :span="20" class="text">{{ transaction.func }}</el-col>
        </el-row>
      </div>
    </el-row>
    <el-row>
      <el-col
        :span="12"
        style="border-right:dotted;border-width:1px;text-aligin:left;"
      >
        <div class="head"><span> 请求参数</span></div>
        <div
          v-for="(arg, index) in transaction.args"
          :key="(index, arg)"
          class="arg_text"
        >
          [{{ index }}]{{ arg }}
        </div>
      </el-col>
      <el-col :span="12">
        <div class="head"><span> 回复数据</span></div>
        <div class="arg_text">回复状态：{{ transaction.resp.status }}</div>
        <div class="arg_text">回复消息：{{ transaction.resp.message }}</div>
        <div class="arg_text">回复数据：{{ transaction.resp.data }}</div>
      </el-col>
    </el-row>
  </el-card>
</template>

<script>
export default {
  props: {
    transaction: Object
  }
};
</script>

<style lang="less">
.head {
  font-weight: bold;
  font-size: small;
  margin: 5px;
}

.arg_text {
  white-space: pre-wrap;
  word-break: break-all;
  text-align: left;
  padding-left: 5px;
}

.text {
  white-space: pre-wrap;
  word-break: break-all;
}

.code_text {
  font-family: Menlo, Monaco, Consolas, Courier, monospace;
  font-size: smaller;
}

.cert_icon {
  margin-left: 6px;
  margin-top: 2px;
}
</style>