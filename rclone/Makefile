PORTNAME=	rclone
DISTVERSIONPREFIX=	v
DISTVERSION=	1.72.0
#PORTREVISION=	1
CATEGORIES=	local net

MAINTAINER=	tremere@cainites.net
COMMENT=	Sync files to and from various cloud services
WWW=		https://rclone.org/

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/COPYING

RUN_DEPENDS=	ca_root_nss>=0:security/ca_root_nss

USES=		go:modules
GO_MODULE=	github.com/${PORTNAME}/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME} \
		share/man/man1/${PORTNAME}.1.gz

JQ_GETVER_EXPR=	.[0].name | gsub("[-a-z]"; "")

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X '${GO_MODULE}/fs.VersionSuffix='"

OPTIONS_DEFINE=	CMOUNT

CMOUNT_DESC=		Enable cmount command (depends on fusefs-libs)
CMOUNT_LIB_DEPENDS=	libfuse.so:filesystems/fusefs-libs
CMOUNT_VARS=		GO_BUILDFLAGS+="-tags cmount"

post-install:
	${INSTALL_MAN} ${WRKSRC}/${PORTNAME}.1 ${STAGEDIR}${PREFIX}/share/man/man1

.include <bsd.port.mk>
