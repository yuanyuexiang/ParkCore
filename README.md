$ solc --abi ParkingLot.sol -o build

$ abigen --abi ParkingLot.json --pkg main --type ParkingLot --out ParkingLot.go


## 自动生成路由文件

bee generate routers

## 生成自动化文档

bee run -gendoc=true -downdoc=true