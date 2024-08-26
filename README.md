# CloudPlatform - 작성중
## Abstract

연구원들은 연구 환경준비에 있어, 반복적으로 양산되어야하는 인프라 환경이 필연적으로 존재합니다. 예를 들어, 군집 드론 자율주행 모듈에 심층강화학습을 하기 위해서는, 자율주행드론 객체를 100여개 이상 준비하여야 하는 경우가 발생합니다. DRL, Vision, NLP 등 다양한 분야를 연구하고 있는 연구원들이 손쉽게 연구환경을 준비하고 양산할 수 있도록 IaaS 와 PaaS 중간 형태로 연구환경을 한번 지정하면 양산된 서비스를 제공할 수 있도록 클라우드 플랫폼을 개발하게 되었습니다. 

교수님이 수업을 진행하다보면 실습수업도 함께 진행하게 되는데, 100명가까이 수강생들을 위한 실습환경을 수작업으로 학부연구생이나 대학원생들이 환경을 준비하는 모습을 보았습니다. 이러한 문제를 해결하기 위해  IaC를 기반으로 클라우드 서비스가 동작할 수 있도록 시스템을 설계하였습니다.
## 목차
1. [소개](#소개)
    - [Keywords](#keywords)
    - [Components](#components)
    - [Installation Guide](#installation-guide)
    - [Features](#features)
3. [용량산정 및 요구량 추정](#용량산정-및-요구량-추정)
4. [설계](#설계)
5. [개발자 소개](#개발자-소개)

## 소개
### Keywords
 ---
 
 **[API Trasaction](https://regular-parsnip-82d.notion.site/Openstack-CloudPlatform-User-API-d31a59a9dd734f2484dbd734c5465b8d?pvs=4)**
 
 수탁/위탁 구조를 기반으로 시스템을 설계하였습니다. 상위 계층으로 부터 받은 요청을 수탁하고 하위 계층으로 다시 위탁하는 구조입니다. 우리나라 정부사업과제가 진행되는 모습을 모방하여 설계하였습니다. 위탁되는 과정이 많아질 수록 Consistency, Integrity, Durability를 확보하기 힘들어집니다. 이를 위해 API를 Transaction으로 관리할 수 있도록 설계하였습니다. 자세한 내용은 링크를 눌러 설계하게된 경위와 솔루션을 작성하였습니다.
 
 ---
 
 **Autonomic Provisioning(개발연구중)**
 
 수요에 따라 자원을 Up 또는 Down Scailing 하는 것을 Auto Provisioning이라고 합니다. 현재 Openstack,K8s,MaaS 세가지 플랫폼을 통해 자원을 제공하고 있어 세 플랫폼을 한정된 물리적 서버에서 상태를 변경해가며 서로 다른 성격의 플랫폼의 노예 자원으로 제공되는 것을 Autonomic Provisionging이라고 표현하였습니다.
 
 ---
 
 **[IaC Rollback](https://velog.io/@ksun4131/%EB%B3%B5%EC%9E%A1%ED%95%9C-%EC%9D%B8%ED%94%84%EB%9D%BC%EA%B5%AC%EC%A1%B0%EB%A5%BC-%EB%A1%A4%EB%B0%B1%ED%95%B4%EC%95%BC%ED%95%9C%EB%8B%A4%EB%A9%B4)**

 연구진 또는 교수님꼐서 생성한 연구환경을 IaC파일(Terraform, Ansible) 로 관리하고 있습니다. 하지만, 시스템 오류로인해 특정시점으로 서비스를 롤백시켜야하는 상황이 온다면 이 IaC파일들을 특정시점으로 복구할수 없어, 데이터베이스와 IaC파일들 간의 일관성이 꺠지게 됩니다. 뿐만아니라, 파일시스템 오류로인해 IaC파일들의 정보가 휘발된다면 시스템 가용성에 심각한 문제를 일으키게됩니다. 그래서 IaC파일들의 존속성을 보장하며, IaC파일들의 버전을 관리하며 특정시점으로 롤백할 수 있도록 Filesytem Managing Middleware 를 제작하게되었습니다. 참고(https://github.com/KwonSunJae/MuscleMemory.git)


 ---


### Components
- **Skin(구 Client)**: 클라우드플랫폼을 웹서비스형태로 제공해주는 컴포넌트입니다. 마치 피부 같습니다.
- **Blood(구 UDCS)**: API Gateway의 성격을 띄지만, 사용자가 요청한 API에 대해  안정성을 위해 부가적인 행동을 수행하는 미들웨어입니다. 마치 적혈구를 운반하는 혈액과 같습니다.
- **Heart(구 SOMS)**: 클라우드플랫폼의 실질적인 비즈니스 서비스를 담당하는 컴포넌트입니다. 마치 심장과 같습니다.
- **MuscleMemory**: 사용자가 생성한 IaC를 관리하고 영원불변하고 특정시점으로 Rollback기능을 제공하는 컴포넌트입니다. 마치 머슬메모리과 같습니다.
- **Heartbeat(구 nonamed)**: 사용자의 API를 기록하고 시스템 통계지표를 보여주는 모니터링 컴포넌트입니다. 마치 인체의 지표인 심박수와 같습니다.

### Installation Guide - 작성중

BLANK


### Features
- **간편한 UI로 연구환경을 준비**: VNC 웹뷰를 통해 시스템환경에 접근해서 시스템을 구축하거나, Ansible 파일을 작성하여 시스템을 구축할수 있습니다. 
- **연구환경 클러스터 배포**: 연구환경이 만약 세트로 2개의 VM 1개의 Baremetal로 구성될 경우에도, 세트를 양산하여 배포할 수 있습니다.
- **지원되는 자원**: Baremetal, Container, VM, Private Network, 공인IP 자원을 제공할 수 있습니다.

## 용량산정 및 요구량 추정

가정
 - 월간 능동 사용자는 DMSLAB 연구원 30명, 컴퓨터공학부 교수님 20명, 강의 실습생 200명 (방학기간은 Down Scailing)
 - 30%의 사용자는 신청한 자원을 Release 하지않는다(AI 학습용 서버). 나머지 70%의 사용자는 신청 자원의 변동성이 크다 (실습용).
 - 사용자 요청의 80%는 Openstack API를 호출한다.
 - 한사람이 클라우드플랫폼에 접속하고 종료할때 까지 발생하는 API는 30회 정도이다.
 - 하루에 접속하는 사람은 평균 20명 정도이다.

추정
 - (CREATE, UPDATE, DELETE) VM API 1개당 추가되는 Database row 수 추정: (Blood 0.5 row) + (Heart 4.5 row) + (Openstack 10row)  =  15row
 - 각 항목 1row당 크기 : Blood 10 byte, Heart 100byte , Openstack 500byte
 - (CREATE, UPDATE, DELETE) VM API 1개당 발생하는 File write 용량 : (Transaction Log)50byte * 3 + (IaC파일)200byte + (Openstack) 500byte + 기타 = 1kb
 - **1년간 시스템이 볼륨 사용 예상** : 20 * 30 * 356 * 1kb  = 약 200 mb
 - **1년간 데이터베이스 용량 사용 예상** : 20 * 30 * 356 * 0.6kb = 약 120mb

가정
 - 사용자 한명은 1개의 자원을 할당한다.
 - 30%의 사용자는 AI 학습을 위한 500GB이상의 SSD 저장용량을 요구한다.
 - 70%의 사용자는 Ubuntu OS를 사용하기 20GB 가량의 저장용량을 요구한다.
   
추정
 - **볼륨 할당 용량 예상** : 250 * 0.3 * 500GB + 250 * 0.7 * 20GB = 37500 + 3500 = 약 40 TB

가정 
 - 사용자 한명은 1개의 자원을 할당한다.
 - 30%의 사용자는 8core 12GB RAM , GPU 1대를 요구한다.
 - 70%의 사용자는 4core 4GB RAM 을 요구한다.

추정 
 - **Core 할당 예상** : 250 * 0.3 * 8 + 250 * 0.7 * 4 = 1300 core
 - **RAM 할당 예상** : 250 * 0.3 * 12 + 250 * 0.7 * 4 = 1600 RAM

Core : vCore = 1:4 비율로 가상화 400 물리 Core , 1600 GB RAM 준비



**16Core , 64GB ,GPU 2대, SSD 1TB,500GB 구성으로 40U, 8RACK 구성 필요**

**1랙의 서버 총 전력 소비** : [ 5{대} * 1600{W} = 8000{W} ]

**8랙의 총 전력 소비**: [ 8{랙} * 8000{W} = 64000{W} ]

**여유 전력을 30% 추가하여 계산**: [ 64000{W} * 1.3 = 83200{W} ]




## 설계
![logical-view System](https://github.com/KwonSunJae/CloudPlatform/blob/docs/docs/cloudplatform-logical.png)
**Logical View Cloudplatform**

![physical-view System](https://github.com/KwonSunJae/CloudPlatform/blob/docs/docs/cloudplatform-physical-view%20(2).png)

**Left: Multiplized Deployment , Right: Minimalized Deployment Physical View Cloudplatform**


## 개발자 소개
- **권순재**: PM, PL, 아키텍처 및 시스템 설계, 전체 개발 
- **장동수**: Blood(구 UDCS) 개발
- **강민희**: Heart(구 SOMS) Container Service 일부 개발
