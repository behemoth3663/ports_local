# Created by: Devon H. O'Dell <devon.odell@gmail.com>
# $FreeBSD: head/lang/go/Makefile 524508 2020-01-28 18:26:16Z jlaffaye $

PORTNAME=	go
PORTVERSION=	1.13.15
CATEGORIES=	local lang
MASTER_SITES?=	https://golang.org/dl/
PKGNAMESUFFIX?=	-${PORTVERSION:R}
DISTNAME?=	go${PORTVERSION}.src

MAINTAINER=	none@FreeBSD.org
COMMENT=	Go programming language

LICENSE=	BSD3CLAUSE

BUILD_DEPENDS=	go:lang/go

USES=		shebangfix
SHEBANG_LANG+=	sh
SHEBANG_FILES=	misc/benchcmp \
		misc/nacl/go_nacl_*_exec \
		misc/wasm/go_js_wasm_exec \
		src/net/http/cgi/testdata/test.cgi
SHEBANG_GLOB=	*.bash *.pl *.sh

sh_OLD_CMD=	/bin/bash "/usr/bin/env bash"
sh_CMD=		${SH}

# Upstream archive contains files with UTF-8 names
EXTRACT_CMD=	${SETENV} LC_ALL=en_US.UTF-8 ${TAR}

.ifndef MASTERDIR
WRKSRC=		${WRKDIR}/go
.endif

ONLY_FOR_ARCHS=	amd64 armv6 armv7 i386

OPTIONS_DEFINE_i386=	GO387
GO387_DESC=	Do not generate code with SSE2 (for old x86 CPU)

.include <bsd.port.pre.mk>

.if ${ARCH} == i386
GOARCH=386
.elif ${ARCH} == "amd64"
GOARCH=amd64
.elif ${ARCH} == armv6 || ${ARCH} == armv7
GOARCH=arm
GOARM=${ARCH:C/armv//}
.else
IGNORE=		unknown arch ${ARCH}
.endif

.ifdef QEMU_EMULATING
IGNORE=		fails to build with qemu-user-static
.endif

.if ${PORT_OPTIONS:MGO387}
GO386=387
.endif

PLIST_SUB+=	opsys_ARCH=${OPSYS:tl}_${GOARCH}

post-patch:
	${REINPLACE_CMD} -e 's|^if ulimit -T|false \&\& &|' ${WRKSRC}/src/run.bash
	${REINPLACE_CMD} -e 's|\$$([[:space:]]*type[[:space:]][^)]*)|""|' ${WRKSRC}/src/make.bash

do-build:
	cd ${WRKSRC}/src && ${SETENV} \
		XDG_CACHE_HOME=${WRKDIR} \
		GOROOT_BOOTSTRAP=${LOCALBASE}/go \
		GO_LDFLAGS="-s -w" \
		CGO_ENABLED=1 \
		GOROOT=${WRKSRC} \
		GOROOT_FINAL=${PREFIX}/${PORTNAME}${PKGNAMESUFFIX} \
		GOBIN= \
		GOOS=${OPSYS:tl} \
		GOARCH=${GOARCH} \
		GO386=${GO386} \
		GOARM=${GOARM} \
		${SH} make.bash -v

do-install:
	cd ${WRKSRC} && \
		${RM} -r .gitattributes .gitignore .github favicon.ico robots.txt \
		pkg/obj pkg/bootstrap pkg/${OPSYS:tl}_${GOARCH_${ARCH}}/cmd
	${CP} -a ${WRKSRC} ${STAGEDIR}${PREFIX}/${PORTNAME}${PKGNAMESUFFIX}
	cd ${STAGEDIR}${PREFIX} && ${FIND} ${PORTNAME}${PKGNAMESUFFIX} \
		-type f ! -name '*.orig' ! -name '*.bak' >> ${TMPPLIST}

do-test:
	cd ${WRKSRC}/src && ${SETENV} \
		GOROOT=${WRKSRC} \
		PATH=${WRKSRC}/bin:${PATH} \
		GOOS=${OPSYS:tl} \
		GOARCH=${GOARCH} \
		GO386=${GO386} \
		GOARM=${GOARM} \
		${SH} run.bash -no-rebuild

.include <bsd.port.post.mk>
