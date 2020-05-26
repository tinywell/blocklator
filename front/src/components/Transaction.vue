<template>
  <el-card style="margin-top:5px;">
    <el-row>
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
              <div>
                <div class="code_text">
                  用户：{{ transaction.signer.cert.cn }}
                </div>
                <div class="code_text">
                  组织：{{ transaction.signer.cert.org }}
                </div>
                <div class="code_text">
                  角色：{{ transaction.signer.cert.ou }}
                </div>
                <div class="code_text">{{ transaction.signer.cert.pem }}</div>
              </div>
              <i class="el-icon-bank-card" slot="reference"></i>
            </el-popover>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="4" class="label">提案请求者签名：</el-col>
          <el-col :span="20" class="text">{{
            transaction.signer.signature
          }}</el-col>
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
          <el-col :span="4" class="label">交易有效性：</el-col>
          <el-col :span="20" class="text">{{ transaction.filter }}</el-col>
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
    <el-row class="border">
      <el-col :span="12">
        <div class="head"><span> 请求参数</span></div>
        <div
          v-for="(arg, index) in transaction.args"
          :key="(index, arg)"
          class="arg_text"
        >
          [{{ index }}] {{ arg }}
        </div>
      </el-col>
      <el-col :span="12">
        <div class="head"><span> 回复数据</span></div>
        <div class="arg_text">回复状态：{{ transaction.resp.status }}</div>
        <div class="arg_text">回复消息：{{ transaction.resp.message }}</div>
        <div class="arg_text">回复数据：{{ transaction.resp.data }}</div>
      </el-col>
    </el-row>
    <el-row>
      <div class="head">背书签名</div>
      <el-row
        v-for="(e, i) in transaction.endorsers"
        :key="(e, i)"
        class="endorse"
        >[{{ i }}]
        <el-popover placement="right-start" trigger="hover">
          <div>
            <div class="code_text">用户：{{ e.cert.cn }}</div>
            <div class="code_text">组织：{{ e.cert.org }}</div>
            <div class="code_text">角色：{{ e.cert.ou }}</div>
            <div class="code_text">{{ e.cert.pem }}</div>
          </div>
          <i class="el-icon-bank-card" slot="reference"></i>
        </el-popover>
        <span>{{ e.signature }}</span>
      </el-row>
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
  text-align: center;
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

.border {
  border: 1px dotted #e9edf0;
  background-color: #e9edf0;
  text-align: left;
  padding-left: 5px;
  padding-right: 5px;
}

.endorse {
  list-style: none;
  text-align: left;
  white-space: pre-wrap;
  word-break: break-all;
}
</style>