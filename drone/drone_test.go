package drone

import "testing"

func Test_drone_Move(t *testing.T) {
	type fields struct {
		id       string
		x        uint32
		y        uint32
		speed    uint32
		minSpeed uint32
		maxSpeed uint32
	}
	type expectedFields struct {
		id       string
		minSpeed uint32
		maxSpeed uint32
	}
	tests := []struct {
		name   string
		fields fields
		expectedFields expectedFields
	}{
		{
			"Test Move ",
			fields{id: "1", x: 100, y: 200, speed: 20, minSpeed: 33, maxSpeed: 40},
			expectedFields{id: "1", minSpeed: 33, maxSpeed: 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &drone{
				id:       tt.fields.id,
				x:        tt.fields.x,
				y:        tt.fields.y,
				speed:    tt.fields.speed,
				minSpeed: tt.fields.minSpeed,
				maxSpeed: tt.fields.maxSpeed,
			}
			d.Move()
			if d.Id() != tt.expectedFields.id {
				t.Errorf("Error Expected Id of %s, got: %s.", tt.expectedFields.id, d.Id())
			}
			if d.GetSpeed() < tt.expectedFields.minSpeed {
				t.Errorf("Speed should not less than minSpeed of %d got %d", tt.expectedFields.minSpeed, d.GetSpeed())
			}
			if d.GetSpeed() > tt.expectedFields.maxSpeed {
				t.Errorf("Speed should not greater than maxSpeed of %d got %d", tt.expectedFields.maxSpeed, d.GetSpeed())
			}
		})
	}
}
