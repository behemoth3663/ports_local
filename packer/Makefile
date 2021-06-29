PORTNAME=	packer
PORTVERSION=	1.7.3
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	brad@facefault.org
COMMENT=	Tool for creating machine images

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/LICENSE

PLIST_FILES=	bin/${PORTNAME}

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

fetch-version: .PHONY get-version

#_OPTIONS_extract=	501:pre-post-extract 750:post-post-extract

pre-post-extract: .PHONY
	${MV} ${WRKDIR}/ionos-enterprise-sdk-go-4.0.2 ${WRKDIR}/profitbricks-sdk-go-4.0.2

post-post-extract: .PHONY
	${MKDIR} ${WRKSRC}/vendor/github.com/hashicorp/go-getter/gcs ${WRKSRC}/vendor/github.com/hashicorp/go-getter/s3
	${LN} -s ${WRKSRC}/vendor/github.com/hashicorp/go-getter/v2/gcs ${WRKSRC}/vendor/github.com/hashicorp/go-getter/gcs/v2
	${LN} -s ${WRKSRC}/vendor/github.com/hashicorp/go-getter/v2/s3 ${WRKSRC}/vendor/github.com/hashicorp/go-getter/s3/v2

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
