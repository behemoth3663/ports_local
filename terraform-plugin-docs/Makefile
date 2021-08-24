PORTNAME=	terraform-plugin-docs
PORTVERSION=	0.4.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Generate and validate Terraform plugin/provider documentation

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

GO_TARGET=	./cmd/tfplugindocs

PLIST_FILES=	bin/tfplugindocs

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
