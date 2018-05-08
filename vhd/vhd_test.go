package vhd

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestVHDExtraHeader_CookieString(t *testing.T) {
	type fields struct {
		Cookie              [8]byte
		DataOffset          [8]byte
		TableOffset         [8]byte
		HeaderVersion       [4]byte
		MaxTableEntries     [4]byte
		BlockSize           [4]byte
		Checksum            [4]byte
		ParentUUID          [16]byte
		ParentTimestamp     [4]byte
		Reserved            [4]byte
		ParentUnicodeName   [512]byte
		ParentLocatorEntry1 [24]byte
		ParentLocatorEntry2 [24]byte
		ParentLocatorEntry3 [24]byte
		ParentLocatorEntry4 [24]byte
		ParentLocatorEntry5 [24]byte
		ParentLocatorEntry6 [24]byte
		ParentLocatorEntry7 [24]byte
		ParentLocatorEntry8 [24]byte
		Reserved2           [256]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := &VHDExtraHeader{
				Cookie:              tt.fields.Cookie,
				DataOffset:          tt.fields.DataOffset,
				TableOffset:         tt.fields.TableOffset,
				HeaderVersion:       tt.fields.HeaderVersion,
				MaxTableEntries:     tt.fields.MaxTableEntries,
				BlockSize:           tt.fields.BlockSize,
				Checksum:            tt.fields.Checksum,
				ParentUUID:          tt.fields.ParentUUID,
				ParentTimestamp:     tt.fields.ParentTimestamp,
				Reserved:            tt.fields.Reserved,
				ParentUnicodeName:   tt.fields.ParentUnicodeName,
				ParentLocatorEntry1: tt.fields.ParentLocatorEntry1,
				ParentLocatorEntry2: tt.fields.ParentLocatorEntry2,
				ParentLocatorEntry3: tt.fields.ParentLocatorEntry3,
				ParentLocatorEntry4: tt.fields.ParentLocatorEntry4,
				ParentLocatorEntry5: tt.fields.ParentLocatorEntry5,
				ParentLocatorEntry6: tt.fields.ParentLocatorEntry6,
				ParentLocatorEntry7: tt.fields.ParentLocatorEntry7,
				ParentLocatorEntry8: tt.fields.ParentLocatorEntry8,
				Reserved2:           tt.fields.Reserved2,
			}
			if got := header.CookieString(); got != tt.want {
				t.Errorf("VHDExtraHeader.CookieString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVHDExtraHeader_addChecksum(t *testing.T) {
	type fields struct {
		Cookie              [8]byte
		DataOffset          [8]byte
		TableOffset         [8]byte
		HeaderVersion       [4]byte
		MaxTableEntries     [4]byte
		BlockSize           [4]byte
		Checksum            [4]byte
		ParentUUID          [16]byte
		ParentTimestamp     [4]byte
		Reserved            [4]byte
		ParentUnicodeName   [512]byte
		ParentLocatorEntry1 [24]byte
		ParentLocatorEntry2 [24]byte
		ParentLocatorEntry3 [24]byte
		ParentLocatorEntry4 [24]byte
		ParentLocatorEntry5 [24]byte
		ParentLocatorEntry6 [24]byte
		ParentLocatorEntry7 [24]byte
		ParentLocatorEntry8 [24]byte
		Reserved2           [256]byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &VHDExtraHeader{
				Cookie:              tt.fields.Cookie,
				DataOffset:          tt.fields.DataOffset,
				TableOffset:         tt.fields.TableOffset,
				HeaderVersion:       tt.fields.HeaderVersion,
				MaxTableEntries:     tt.fields.MaxTableEntries,
				BlockSize:           tt.fields.BlockSize,
				Checksum:            tt.fields.Checksum,
				ParentUUID:          tt.fields.ParentUUID,
				ParentTimestamp:     tt.fields.ParentTimestamp,
				Reserved:            tt.fields.Reserved,
				ParentUnicodeName:   tt.fields.ParentUnicodeName,
				ParentLocatorEntry1: tt.fields.ParentLocatorEntry1,
				ParentLocatorEntry2: tt.fields.ParentLocatorEntry2,
				ParentLocatorEntry3: tt.fields.ParentLocatorEntry3,
				ParentLocatorEntry4: tt.fields.ParentLocatorEntry4,
				ParentLocatorEntry5: tt.fields.ParentLocatorEntry5,
				ParentLocatorEntry6: tt.fields.ParentLocatorEntry6,
				ParentLocatorEntry7: tt.fields.ParentLocatorEntry7,
				ParentLocatorEntry8: tt.fields.ParentLocatorEntry8,
				Reserved2:           tt.fields.Reserved2,
			}
			h.addChecksum()
		})
	}
}

func TestVHDHeader_DiskTypeStr(t *testing.T) {
	type fields struct {
		Cookie             [8]byte
		Features           [4]byte
		FileFormatVersion  [4]byte
		DataOffset         [8]byte
		Timestamp          [4]byte
		CreatorApplication [4]byte
		CreatorVersion     [4]byte
		CreatorHostOS      [4]byte
		OriginalSize       [8]byte
		CurrentSize        [8]byte
		DiskGeometry       [4]byte
		DiskType           [4]byte
		Checksum           [4]byte
		UniqueId           [16]byte
		SavedState         [1]byte
		Reserved           [427]byte
	}
	tests := []struct {
		name   string
		fields fields
		wantDt string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &VHDHeader{
				Cookie:             tt.fields.Cookie,
				Features:           tt.fields.Features,
				FileFormatVersion:  tt.fields.FileFormatVersion,
				DataOffset:         tt.fields.DataOffset,
				Timestamp:          tt.fields.Timestamp,
				CreatorApplication: tt.fields.CreatorApplication,
				CreatorVersion:     tt.fields.CreatorVersion,
				CreatorHostOS:      tt.fields.CreatorHostOS,
				OriginalSize:       tt.fields.OriginalSize,
				CurrentSize:        tt.fields.CurrentSize,
				DiskGeometry:       tt.fields.DiskGeometry,
				DiskType:           tt.fields.DiskType,
				Checksum:           tt.fields.Checksum,
				UniqueId:           tt.fields.UniqueId,
				SavedState:         tt.fields.SavedState,
				Reserved:           tt.fields.Reserved,
			}
			if gotDt := h.DiskTypeStr(); gotDt != tt.wantDt {
				t.Errorf("VHDHeader.DiskTypeStr() = %v, want %v", gotDt, tt.wantDt)
			}
		})
	}
}

func TestVHDHeader_TimestampTime(t *testing.T) {
	type fields struct {
		Cookie             [8]byte
		Features           [4]byte
		FileFormatVersion  [4]byte
		DataOffset         [8]byte
		Timestamp          [4]byte
		CreatorApplication [4]byte
		CreatorVersion     [4]byte
		CreatorHostOS      [4]byte
		OriginalSize       [8]byte
		CurrentSize        [8]byte
		DiskGeometry       [4]byte
		DiskType           [4]byte
		Checksum           [4]byte
		UniqueId           [16]byte
		SavedState         [1]byte
		Reserved           [427]byte
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &VHDHeader{
				Cookie:             tt.fields.Cookie,
				Features:           tt.fields.Features,
				FileFormatVersion:  tt.fields.FileFormatVersion,
				DataOffset:         tt.fields.DataOffset,
				Timestamp:          tt.fields.Timestamp,
				CreatorApplication: tt.fields.CreatorApplication,
				CreatorVersion:     tt.fields.CreatorVersion,
				CreatorHostOS:      tt.fields.CreatorHostOS,
				OriginalSize:       tt.fields.OriginalSize,
				CurrentSize:        tt.fields.CurrentSize,
				DiskGeometry:       tt.fields.DiskGeometry,
				DiskType:           tt.fields.DiskType,
				Checksum:           tt.fields.Checksum,
				UniqueId:           tt.fields.UniqueId,
				SavedState:         tt.fields.SavedState,
				Reserved:           tt.fields.Reserved,
			}
			if got := h.TimestampTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VHDHeader.TimestampTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVHDHeader_addChecksum(t *testing.T) {
	type fields struct {
		Cookie             [8]byte
		Features           [4]byte
		FileFormatVersion  [4]byte
		DataOffset         [8]byte
		Timestamp          [4]byte
		CreatorApplication [4]byte
		CreatorVersion     [4]byte
		CreatorHostOS      [4]byte
		OriginalSize       [8]byte
		CurrentSize        [8]byte
		DiskGeometry       [4]byte
		DiskType           [4]byte
		Checksum           [4]byte
		UniqueId           [16]byte
		SavedState         [1]byte
		Reserved           [427]byte
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &VHDHeader{
				Cookie:             tt.fields.Cookie,
				Features:           tt.fields.Features,
				FileFormatVersion:  tt.fields.FileFormatVersion,
				DataOffset:         tt.fields.DataOffset,
				Timestamp:          tt.fields.Timestamp,
				CreatorApplication: tt.fields.CreatorApplication,
				CreatorVersion:     tt.fields.CreatorVersion,
				CreatorHostOS:      tt.fields.CreatorHostOS,
				OriginalSize:       tt.fields.OriginalSize,
				CurrentSize:        tt.fields.CurrentSize,
				DiskGeometry:       tt.fields.DiskGeometry,
				DiskType:           tt.fields.DiskType,
				Checksum:           tt.fields.Checksum,
				UniqueId:           tt.fields.UniqueId,
				SavedState:         tt.fields.SavedState,
				Reserved:           tt.fields.Reserved,
			}
			h.addChecksum()
		})
	}
}

func Test_calculateCHS(t *testing.T) {
	type args struct {
		ts uint64
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCHS(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateCHS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexToField(t *testing.T) {
	type args struct {
		hexs  string
		field []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hexToField(tt.args.hexs, tt.args.field)
		})
	}
}

func Test_getMaxTableEntries(t *testing.T) {
	type args struct {
		diskSize uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxTableEntries(tt.args.diskSize); got != tt.want {
				t.Errorf("getMaxTableEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateFixedHeader(t *testing.T) {
	type args struct {
		size    uint64
		options *VHDOptions
	}
	y , _ := CreateFixedHeader(12, &VHDOptions{
		UUID: "38323065-3932-3030-2d61-3564322d3764",
	})
	tests := []struct {
		name     string
		args     args
		wantVhdh *VHDHeader
		wantErr  bool
	}{
		{
			name: "Default",
			args: args{
				size: 12,
				options: &VHDOptions{
					UUID: "38323065-3932-3030-2d61-3564322d3764",
				},
			},
			wantErr: false,
			wantVhdh: y,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVhdh, err := CreateFixedHeader(tt.args.size, tt.args.options)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFixedHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVhdh, tt.wantVhdh) {
				t.Errorf("CreateFixedHeader() = %v, want %v", gotVhdh, tt.wantVhdh)
			}
		})
	}
}

func TestRawToFixed(t *testing.T) {
	type args struct {
		f       *os.File
		options *VHDOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RawToFixed(tt.args.f, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("RawToFixed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVHDCreateSparse(t *testing.T) {
	type args struct {
		size    uint64
		name    string
		options VHDOptions
	}
	tests := []struct {
		name    string
		args    args
		wantVhd *VHD
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVhd, err := VHDCreateSparse(tt.args.size, tt.args.name, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("VHDCreateSparse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVhd, tt.wantVhd) {
				t.Errorf("VHDCreateSparse() = %v, want %v", gotVhd, tt.wantVhd)
			}
		})
	}
}

func TestFromFile(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name    string
		args    args
		wantVhd *VHD
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVhd, err := FromFile(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVhd, tt.wantVhd) {
				t.Errorf("FromFile() = %v, want %v", gotVhd, tt.wantVhd)
			}
		})
	}
}

func TestVHD_PrintInfo(t *testing.T) {
	type fields struct {
		Footer      *VHDHeader
		ExtraHeader *VHDExtraHeader
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vhd := &VHD{
				Footer:      tt.fields.Footer,
				ExtraHeader: tt.fields.ExtraHeader,
			}
			vhd.PrintInfo()
		})
	}
}

func TestVHD_PrintExtraHeader(t *testing.T) {
	type fields struct {
		Footer      *VHDHeader
		ExtraHeader *VHDExtraHeader
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vhd := &VHD{
				Footer:      tt.fields.Footer,
				ExtraHeader: tt.fields.ExtraHeader,
			}
			vhd.PrintExtraHeader()
		})
	}
}

func TestVHD_PrintFooter(t *testing.T) {
	type fields struct {
		Footer      *VHDHeader
		ExtraHeader *VHDExtraHeader
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vhd := &VHD{
				Footer:      tt.fields.Footer,
				ExtraHeader: tt.fields.ExtraHeader,
			}
			vhd.PrintFooter()
		})
	}
}

func Test_readVHDExtraHeader(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name       string
		args       args
		wantHeader *VHDExtraHeader
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, err := readVHDExtraHeader(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("readVHDExtraHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("readVHDExtraHeader() = %v, want %v", gotHeader, tt.wantHeader)
			}
		})
	}
}

func Test_readVHDFooter(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name       string
		args       args
		wantHeader *VHDHeader
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, err := readVHDFooter(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("readVHDFooter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("readVHDFooter() = %v, want %v", gotHeader, tt.wantHeader)
			}
		})
	}
}

func Test_readVHDHeader(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name       string
		args       args
		wantHeader *VHDHeader
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, err := readVHDHeader(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("readVHDHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("readVHDHeader() = %v, want %v", gotHeader, tt.wantHeader)
			}
		})
	}
}
