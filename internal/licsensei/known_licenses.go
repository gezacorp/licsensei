// SPDX-License-Identifier: MIT
// Copyright (c) 2025 Geza Corp authors
package licsensei

const (
	LicenseApache20 = `Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.`
	LicenseApache20SPDX = "Apache-2.0"

	LicenseMIT = `Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.`
	LicenseMITSPDX = "MIT"

	LicenseBSD = `Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.`
	LicenseBSD2SPDX = "BSD-2-Clause"
	LicenseBSD3SPDX = "BSD-3-Clause"

	LicenseMPL2 = `This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.`
	LicenseMPL2SPDX = "MPL-2.0"

	LicenseGPL20Only = `This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; version 2.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, see
<https://www.gnu.org/licenses/>.`
	LicenseGPL20OnlySPDX = "GPL-2.0-only"

	LicenseGPL20OrLater = `This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, see
<https://www.gnu.org/licenses/>.`
	LicenseGPL20OrLaterSPDX = "GPL-2.0-or-later"

	LicenseGPL30Only = ` This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, version 3.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.`
	LicenseGPL30OnlySPDX = "GPL-3.0-only"

	LicenseGPL30OrLater = ` This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.`
	LicenseGPL30OrLaterSPDX = "GPL-3.0-or-later"
)

var knownLicenses = map[string]string{
	LicenseApache20SPDX:     LicenseApache20,
	LicenseBSD2SPDX:         LicenseBSD,
	LicenseBSD3SPDX:         LicenseBSD,
	LicenseGPL20OnlySPDX:    LicenseGPL20Only,
	LicenseGPL20OrLaterSPDX: LicenseGPL20OrLater,
	LicenseGPL30OnlySPDX:    LicenseGPL30Only,
	LicenseGPL30OrLaterSPDX: LicenseGPL30OrLater,
	LicenseMITSPDX:          LicenseMIT,
	LicenseMPL2SPDX:         LicenseMPL2,
}
