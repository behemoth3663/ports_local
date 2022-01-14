# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	eksctl
PORTVERSION=	0.79.0
CATEGORIES=	local sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	CLI for Amazon EKS

LICENSE=	APACHE20

RUN_DEPENDS=	kubectl:sysutils/kubectl \
		aws-iam-authenticator:security/aws-iam-authenticator

USES=		go:modules

PORTSCOUT=	limit:^[0-9\.]*$$

USE_GITHUB=	yes
GH_ACCOUNT=	weaveworks

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}
GO_BUILDFLAGS=	-ldflags="-X github.com/weaveworks/eksctl/pkg/version.gitCommit=${PORTVERSION} \
		-X github.com/weaveworks/eksctl/pkg/version.builtAt=${%s:L:gmtime}"

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
