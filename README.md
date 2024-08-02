# CloudPlatform

## 목차
1. [소개](#소개)
    - [Keywords](#keywords)
    - [Abstract](#abstract)
    - [Components](#components)
    - [Installation Guide](#installation-guide)
    - [Features](#features)
3. [요구사항](#요구사항)
4. [설계](#설계)
5. [개발자 소개](#개발자-소개)

## 소개
### Keywords
 ---
 
 **[API Trasaction](https://regular-parsnip-82d.notion.site/Openstack-CloudPlatform-User-API-d31a59a9dd734f2484dbd734c5465b8d?pvs=4)**
 
 수탁/위탁 구조를 기반으로 시스템을 설계하였습니다. 상위 계층으로 부터 받은 요청을 수탁하고 하위 계층으로 다시 위탁하는 구조입니다. 우리나라 정부사업과제가 진행되는 모습을 모방하여 설계하였습니다. 위탁되는 과정이 많아질 수록 Consistency, Integrity, Durability를 확보하기 힘들어집니다. 이를 위해 API를 Transaction으로 관리할 수 있도록 설계하였습니다.
 
 ---
 
 **Autonomic Provisioning**
 
 수요에 따라 자원을 Up 또는 Down Scailing 하는 것을 Auto Provisioning이라고 합니다. 현재 Openstack,K8s,MaaS 세가지 플랫폼을 통해 자원을 제공하고 있어 세 플랫폼을 한정된 물리적 서버에서 상태를 변경해가며 서로 다른 성격의 플랫폼의 노예 자원으로 제공되는 것을 Autonomic Provisionging이라고 표현하였습니다.
 
 ---
 
 **[IaC Rollback](https://velog.io/@ksun4131/%EB%B3%B5%EC%9E%A1%ED%95%9C-%EC%9D%B8%ED%94%84%EB%9D%BC%EA%B5%AC%EC%A1%B0%EB%A5%BC-%EB%A1%A4%EB%B0%B1%ED%95%B4%EC%95%BC%ED%95%9C%EB%8B%A4%EB%A9%B4)**
 
 ---

### Abstract
개발한 클라우드플랫폼은 Openstack,K8s,MaaS가 제공하는 서비스들을 추상화하여 IaaS를 SaaS로 제공할 수 있도록 하고자 합니다. 모든 자원들을 IaC로 관리하여 연구진들의 연구환경, 수업을 진행할 떄의 실습환경들을 빠르게 양산하고 배포할 수 있도록 합니다.


### Components
- **Skin(구 Client)**: 클라우드플랫폼을 웹서비스형태로 제공해주는 컴포넌트입니다. 마치 피부 같습니다.
- **Blood(구 UDCS)**: API Gateway의 성격을 띄지만, 사용자가 요청한 API에 대해  안정성을 위해 부가적인 행동을 수행하는 미들웨어입니다. 마치 적혈구를 운반하는 혈액과 같습니다.
- **Heart(구 SOMS)**: 클라우드플랫폼의 실질적인 비즈니스 서비스를 담당하는 컴포넌트입니다. 마치 심장과 같습니다.
- **MuscleMemory**: 사용자가 생성한 IaC를 관리하고 영원불변하고 특정시점으로 Rollback기능을 제공하는 컴포넌트입니다. 마치 머슬메모리과 같습니다.
- **Heartbeat(구 nonamed)**: 사용자의 API를 기록하고 시스템 통계지표를 보여주는 모니터링 컴포넌트입니다. 마치 인체의 지표인 심박수와 같습니다.

### Installation Guide


### Features
- **Feature 1**: 설명
- **Feature 2**: 설명
- **Feature 3**: 설명

## 요구사항


## 설계
![logical-view System](https://github.com/KwonSunJae/CloudPlatform/blob/docs/docs/cloudplatform-logical.png)
**Logical View Cloudplatform**

![physical-view System] (https://github.com/KwonSunJae/CloudPlatform/blob/docs/docs/cloudplatform-physical-view%20(2).png)

**Left: Multiplized Deployment , Right: Minimalized Deployment Physical Architecture**


## 개발자 소개
- **권순재**: PM, PL, 아키텍처 및 시스템 설계, 전체 개발 
- **장동수**: Blood(구 UDCS) 개발
- **강민희**: Heart(구 SOMS) Container Service 일부 개발
