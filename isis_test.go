package isis

import (
	"net/netip"
	"os/exec"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidate(t *testing.T) {
	type args struct {
		ipAdrr string
	}
	tests := []struct {
		args args
		want netip.Addr
	}{
		{
			args: args{ipAdrr: "10.0.0.0"},
			want: netip.MustParseAddr("10.0.0.0"),
		},
		{
			args: args{ipAdrr: "172.16.255.1"},
			want: netip.MustParseAddr("172.16.255.1"),
		},
		{
			args: args{ipAdrr: "192.168.20.40"},
			want: netip.MustParseAddr("192.168.20.40"),
		},
		{
			args: args{ipAdrr: "0.0.0.0"},
			want: netip.MustParseAddr("0.0.0.0"),
		},
		{
			args: args{ipAdrr: "255.255.255.255"},
			want: netip.MustParseAddr("255.255.255.255"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.ipAdrr, func(t *testing.T) {
			if got := Validate(tt.args.ipAdrr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToNET(t *testing.T) {
	type args struct {
		ipAdrr netip.Addr
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{ ipAdrr: netip.MustParseAddr("1.1.1.1")},
			want: "0010.0100.1001",
		},
		{
			args: args{ ipAdrr: netip.MustParseAddr("2.2.2.2")},
			want: "0020.0200.2002",
		},
		{
			args: args{ ipAdrr: netip.MustParseAddr("10.10.10.10")},
			want: "0100.1001.0010",
		},
		{
			args: args{ ipAdrr: netip.MustParseAddr("192.168.255.1")},
			want: "1921.6825.5001",
		},
		{
			args: args{ ipAdrr: netip.MustParseAddr("255.255.255.255")},
			want: "2552.5525.5255",
		},
		{
			args: args{ ipAdrr: netip.MustParseAddr("10.128.104.16")},
			want: "0101.2810.4016",
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.ipAdrr.String(), func(t *testing.T) {
			if got := ConvertToNET(tt.args.ipAdrr); got != tt.want {
				t.Errorf("ConvertToNET() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(t *testing.T) {
	t.Parallel()

	want := "NET: 49.0051.0010.0100.1001.00"
	cmd := exec.Command("./ipv4tonet", "-area-prefix", "49", "-area", "0051", "-ip", "1.1.1.1")

	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	} else {
		if !cmp.Equal(want, string(got)) {
			t.Errorf("\nwant: \n%s \ngot: \n%s", want, string(got))
		}

	}
}
