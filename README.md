# Monolith_demo
The monolith and microservice designs have been discussed in many papers. So I omit discussion of it here. I think monolith is a better choice for entrepreneurial projects.
Monolith_demo project is an example of the [elegant monolith or neomonolith](https://inconshreveable.com/10-07-2015/the-neomonolith/), that build by the [Go-kit toolkit](https://github.com/go-kit/kit) with the [Gokit-cli](https://github.com/GrantZheng/kit). It describes a live stream application consisting of multiple services, such as room, gift, user and so on.

# How to build an easily separated monolith

## Software design and modeling methodology
To construct a elegant application，software analyze and modeling is an indispensable procedure when you meet up with a little complex business scenarios. As we all know，DDD(Domain Driven Design) is a very excellent principle put forward by the Eric Evans，I design and build the demo project based on UCDD(Use Case Driven Desgin).

## Elegant implementation based on the modeling
According to the UCDD principle, we can get our business model as shown below, described by the [layered architecture pattern](https://www.oreilly.com/library/view/software-architecture-patterns/9781491971437/ch01.html). It is laid out in three layers: the presentation layer, the business service layer and the basic service layer. The presentation layer is not in the scope of our discussion, the upper layer is dependent on the lower layer. Each layer consist of serveral services.
About the service design, we'd better forget the traditional idea of MVC in our mind, everything is service. The simple live stream application consist of twelve services, each service module corresponds to a folder which generated by the [Go-kit toolkit](https://github.com/go-kit/kit) with the [Gokit-cli](https://github.com/GrantZheng/kit).
![image](https://github.com/GrantZheng/monolith_demo/blob/main/images/archeture.png) 

### About the business service layer
- barrage service: it is dependent on the room service, the user service, the sensitive words service and the message channel service，implementing the barrage business logic.
- living service: it is dependent on the room service, the user service, the video content service and the identity service, implementing the living business logic.
- gift service: it is dependent on the room service, the user service, the identity service, the pay service and the message channel service, implementing the gift business logic.
### About the basic service layer
- user service: it implements the user information logic, such as managing user identity and so on.
- room service: it implements the room information logic.
- wallet service: it implements the wallet logic, such as the account balance.
- pay service: it implements the room pay logic, such as payment channels and reconciliation.
- video_content service: it implements the video content logic, such as managing video meta data, generating play links.
- goods service: it implements the goods logic, such as managing goods meta data.
- message_channel service: it implements the message channel logic, ensuring reliable message delivery.

### About the communication between services
The communication between services must be completed through the interface defined in the service-name/pkg/service/service.go.It is a very critical point for building a monolith that is easy to split.

## Usage
```bash
git clone https://github.com/GrantZheng/monolith_demo.git
cd monolith_demo
go run main.go
curl -X POST -d '{"from":"a","to":"b"}' "http://127.0.0.1:8081/send"
```







