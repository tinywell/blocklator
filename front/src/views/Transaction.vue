<template>
  <div style="display: flex;justify-content: center;">
    <div style="width:900px">
      <blockhead :block="block"></blockhead>
      <transaction
        v-for="tran in block.transactions"
        :key="tran"
        :transaction="tran"
      ></transaction>
    </div>
  </div>
</template>

<script>
import transaction from "../components/Transaction.vue";
import blockhead from "../components/BlockHead.vue";

export default {
  components: { transaction, blockhead },
  data() {
    return {
      block: {}
      // block: {
      //   block_num: 5,
      //   hash:
      //     "6C27910CF97DD41DF1C633214B798450E64ACE662DF47A0CA9765A250C98D7F6",
      //   pre_hash:
      //     "ED0A162888DA86361358D61870564642602CF7DEB718199029F282C3FCEA6ED6",
      //   channel: "mychannel",
      //   type: 0,
      //   transactions: [
      //     {
      //       channel: "mychannel",
      //       tx_id:
      //         "6a77bb2275e7b257ba4d2e0b5ead4d2db34a8940eb23b51a3eec3b5a079f565c",
      //       time: "2020-05-20T11:28:54.763+08:00",
      //       chaincode: "mycc",
      //       func: "PutKey",
      //       args: [
      //         "Org1MSP",
      //         "Admin@org1.example.com",
      //         '{"UserInfo":{"cert":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEcTEbSyRxJBziSkcNnydW92QaM3tXbNfciPRG6R+R\\r\\n1byoceBn/cDNL615IeGUhxRp/35XnK5IuPjk1D5SrhLLhg==","iD":0,"name":"Test7","regCode":"GU9L4E"}}'
      //       ],
      //       resp: {
      //         status: 200,
      //         message: ""
      //       },
      //       signer: {
      //         mspid: "",
      //         cert: ""
      //       },
      //       signature: ""
      //     }
      //   ]
      // }
    };
  },
  methods: {
    init() {
      if (localStorage.block) {
        this.block = JSON.parse(localStorage.block);
      } else {
        this.$message({
          message: "没有区块数据",
          type: "warning"
        });
      }
    }
  },
  mounted() {
    this.init();
  }
};
</script>

<style>
</style>