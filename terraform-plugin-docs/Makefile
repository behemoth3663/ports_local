#

PORTNAME=	terraform-plugin-docs
PORTVERSION=	0.14.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Generate and validate Terraform plugin/provider documentation
WWW=		https://github.com/hashicorp/terraform-plugin-docs

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

GO_TARGET=	./cmd/tfplugindocs

PLIST_FILES=	bin/tfplugindocs

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
