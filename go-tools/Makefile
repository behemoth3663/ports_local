PORTNAME=	tools
DISTVERSIONPREFIX=	v
DISTVERSION=	0.39.0
PORTREVISION=	1
PORTEPOCH=	1
CATEGORIES=	local devel
PKGNAMEPREFIX=	go-

MAINTAINER=	fuz@FreeBSD.org
COMMENT=	Go Tools
WWW=		https://godoc.org/golang.org/x/tools

LICENSE=	BSD3CLAUSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	golang
GH_PROJECT=	${PORTNAME}

# find $(make -V WRKSRC)/cmd -name '*.go' -type f | sed -r -e '/\/testdata\//d' -e '/\/internal\//d' -e 's,.*/cmd/,\t./cmd/,' -e 's,/[^/]+$, \\,' | sort --unique
GO_TARGET=	\
	./cmd/auth/authtest \
	./cmd/auth/cookieauth \
	./cmd/auth/gitauth \
	./cmd/auth/netrcauth \
	./cmd/benchcmp \
	./cmd/bisect \
	./cmd/bundle \
	./cmd/callgraph \
	./cmd/compilebench \
	./cmd/deadcode \
	./cmd/digraph \
	./cmd/eg \
	./cmd/file2fuzz \
	./cmd/fiximports \
	./cmd/go-contrib-init \
	./cmd/godex \
	./cmd/goimports \
	./cmd/gomvpkg \
	./cmd/gonew \
	./cmd/gotype \
	./cmd/goyacc \
	./cmd/html2article \
	./cmd/present \
	./cmd/present2md \
	./cmd/signature-fuzzer/fuzz-driver \
	./cmd/signature-fuzzer/fuzz-runner \
	./cmd/splitdwarf \
	./cmd/ssadump \
	./cmd/stress \
	./cmd/stringer \
	./cmd/toolstash

CONFLICTS_INSTALL=	easygit rubygem-bundler

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.version=${PORTVERSION}'"

PLIST_FILES=	${GO_TARGET:T:S,^,bin/,}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${FIND} ${WRKSRC}/cmd -name go.mod -type f -delete

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-golang:telemetry:bda55230c457:golang_telemetry/vendor/golang.org/x/telemetry
GH_TUPLE:=	${GH_TUPLE:Ngolang\:telemetry\:*}

.include <bsd.port.mk>
