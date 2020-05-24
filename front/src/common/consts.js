const ServerAddrPre = 'http://localhost:8080/api';

const BlockTypeOpts = [
    { text: '配置区块', value: 1 },
    { text: '交易区块', value: 0 },
  ];
  
const BlockTypeOptsMap = {};
BlockTypeOpts.forEach((item) => { BlockTypeOptsMap[item.value] = item.text; });
  

export default {
    BlockTypeOpts,
    BlockTypeOptsMap,
    ServerAddrPre
}