<template>
  <el-card style="margin-top: 5px">
    <el-row>
      <div style="margin-bottom: 5px">
        <el-row>
          <el-col :span="4" class="label">提案请求MSPID：</el-col>
          <el-col :span="20" class="text">{{
            transaction.signer.mspid
          }}</el-col>
        </el-row>
        <!-- <el-row>
          <el-col :span="4" class="label">提案请求者证书：</el-col>
          <el-col :span="20" class="text">
            <el-popover placement="right-start" trigger="hover">
              <cert :cert="transaction.signer.cert"></cert>
              <i class="el-icon-bank-card" slot="reference"></i>
            </el-popover>
          </el-col>
        </el-row> -->
        <el-row>
          <el-col :span="4" class="label">提案请求者签名：</el-col>
          <el-col :span="20" class="text">
            <el-popover placement="right-start" trigger="hover">
              <cert :cert="transaction.signer.cert"></cert>
              <i class="el-icon-bank-card" slot="reference"></i>
            </el-popover>
            {{ transaction.signer.signature }}</el-col
          >
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
          <el-col :span="20" class="text"
            >{{ transaction.filter }} | {{ transaction.validation_code }}
          </el-col>
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
      <el-row
        class="border"
        v-for="rwset in transaction.tx_rw_set.ns_rw_sets"
        :key="rwset"
      >
        <el-row class="head">读写集 namespace: {{ rwset.name_space }}</el-row>
        <el-row>
          <el-col :span="12">
            <div class="head" style="text-align: center">读集</div>
            <div
              class="arg_text"
              v-for="read in rwset.kv_rw_set.reads"
              :key="read"
            >
              <div class="arg_text"><span>key:</span> {{ read.key }}</div>
              <div class="arg_text" v-if="read.version">
                <span>version: </span> {block:{{
                  read.version.block_num
                }}
                txnum:{{ read.version.tx_num }}
                }
              </div>
            </div>
          </el-col>
          <el-col :span="12">
            <div class="head" style="text-align: center">写集</div>
            <div
              class="arg_text"
              v-for="write in rwset.kv_rw_set.writes"
              :key="write"
            >
              <div class="arg_text"><span>key:</span> {{ write.key }}</div>
              <div class="arg_text"><span>value: </span>{{ write.value }}</div>
            </div>
          </el-col>
        </el-row>
      </el-row>
    </el-row>
    <el-row>
      <div class="head">背书签名</div>
      <el-row
        v-for="(e, i) in transaction.endorsers"
        :key="(e, i)"
        class="endorse"
        >[{{ i }}]
        <el-popover placement="right-start" trigger="hover">
          <cert :cert="e.cert"></cert>
          <i class="el-icon-bank-card" slot="reference"></i>
        </el-popover>
        <span>{{ e.signature }}</span>
      </el-row>
    </el-row>
  </el-card>
</template>

<script>
import cert from "./Cert.vue";

export default {
  components: {
    cert
  },
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
  margin-top: 2px;
}

.endorse {
  list-style: none;
  text-align: left;
  white-space: pre-wrap;
  word-break: break-all;
}
</style>