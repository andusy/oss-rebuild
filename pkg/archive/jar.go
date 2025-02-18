// Copyright 2025 Google LLC
// SPDX-License-Identifier: Apache-2.0

package archive

import (
	"bytes"
	"strings"
)

var StableJARBuildMetadata = ZipEntryStabilizer{
	Name: "jar-build-metadata",
	Func: func(zf *MutableZipFile) {
		// Only process MANIFEST.MF files
		if !strings.HasSuffix(zf.Name, "META-INF/MANIFEST.MF") {
			return
		}
		r, err := zf.Open()
		if err != nil {
			return
		}
		manifest, err := ParseManifest(r)
		if err != nil {
			return
		}
		for _, attr := range []string{
			"Archiver-Version",
			"Bnd-LastModified",
			"Build-Jdk",
			"Build-Jdk-Spec",
			"Build-Number",
			"Build-Time",
			"Built-By",
			"Built-Date",
			"Built-Host",
			"Built-OS",
			"Created-By",
			"Hudson-Build-Number",
			"Implementation-Build-Date",
			"Implementation-Build-Java-Vendor",
			"Implementation-Build-Java-Version",
			"Implementation-Build",
			"Jenkins-Build-Number",
			"Originally-Created-By",
			"Os-Version",
			"SCM-Git-Branch",
			"SCM-Revision",
			"SCM-Git-Commit-Dirty",
			"SCM-Git-Commit-ID",
			"SCM-Git-Commit-Abbrev",
			"SCM-Git-Commit-Description",
			"SCM-Git-Commit-Timestamp",
			"Source-Date-Epoch",
		} {
			manifest.MainSection.Delete(attr)
		}
		buf := bytes.NewBuffer(nil)
		if err := WriteManifest(buf, manifest); err != nil {
			return
		}
		zf.SetContent(buf.Bytes())
	},
}
