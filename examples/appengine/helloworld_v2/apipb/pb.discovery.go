// Code generated by cproto. DO NOT EDIT.

package apipb

import discovery "go.chromium.org/luci/grpc/discovery"

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"luci.examples.k8s.helloworld.Greeter",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 164, 86, 81, 111, 219, 200,
			17, 166, 37, 199, 113, 55, 233, 93, 160, 22, 113, 145, 22, 185,
			57, 55, 215, 36, 173, 66, 25, 78, 31, 210, 4, 40, 64, 137,
			107, 107, 115, 52, 169, 146, 148, 125, 190, 167, 80, 228, 74, 218,
			30, 201, 101, 119, 151, 118, 132, 195, 253, 153, 62, 246, 23, 180,
			143, 5, 250, 208, 135, 2, 253, 41, 125, 45, 80, 160, 40, 118,
			73, 217, 178, 15, 105, 81, 156, 94, 180, 156, 153, 157, 157, 249,
			118, 118, 190, 65, 127, 236, 161, 31, 47, 56, 95, 228, 116, 80,
			9, 174, 248, 172, 158, 15, 104, 81, 169, 149, 109, 62, 123, 31,
			55, 74, 123, 173, 220, 191, 139, 238, 96, 173, 31, 94, 160, 31,
			164, 188, 176, 111, 233, 135, 200, 104, 39, 250, 115, 178, 245, 229,
			211, 5, 83, 203, 122, 102, 167, 188, 24, 44, 120, 158, 148, 139,
			235, 99, 42, 181, 170, 168, 108, 78, 251, 231, 214, 214, 239, 59,
			221, 227, 201, 240, 15, 157, 199, 199, 141, 199, 73, 107, 103, 159,
			209, 60, 255, 188, 228, 151, 101, 172, 237, 223, 254, 251, 1, 218,
			233, 109, 63, 182, 94, 62, 64, 127, 187, 143, 182, 238, 247, 186,
			143, 173, 222, 225, 159, 239, 131, 217, 144, 242, 28, 134, 245, 124,
			78, 133, 132, 23, 208, 184, 122, 42, 33, 75, 84, 2, 172, 84,
			84, 164, 203, 164, 92, 80, 152, 115, 81, 36, 10, 193, 136, 87,
			43, 193, 22, 75, 5, 135, 7, 7, 175, 218, 13, 64, 202, 212,
			6, 112, 242, 28, 140, 78, 130, 160, 146, 138, 11, 154, 217, 8,
			150, 74, 85, 242, 245, 96, 144, 209, 11, 154, 243, 138, 10, 185,
			198, 64, 39, 89, 181, 65, 188, 152, 53, 65, 12, 16, 130, 144,
			102, 76, 42, 193, 102, 181, 98, 188, 132, 164, 204, 160, 150, 20,
			88, 9, 146, 215, 34, 165, 70, 50, 99, 101, 34, 86, 38, 46,
			217, 135, 75, 166, 150, 192, 133, 249, 231, 181, 66, 80, 240, 140,
			205, 89, 154, 104, 15, 125, 72, 4, 133, 138, 138, 130, 41, 69,
			51, 168, 4, 191, 96, 25, 205, 64, 45, 19, 5, 106, 169, 179,
			203, 115, 126, 201, 202, 5, 164, 188, 204, 152, 222, 36, 245, 38,
			4, 5, 85, 175, 17, 2, 253, 251, 249, 173, 192, 36, 240, 249,
			58, 162, 148, 103, 20, 138, 90, 42, 16, 84, 37, 172, 52, 94,
			147, 25, 191, 208, 170, 22, 49, 4, 37, 87, 44, 165, 125, 80,
			75, 38, 33, 103, 82, 105, 15, 155, 39, 150, 217, 173, 112, 50,
			38, 211, 60, 97, 5, 21, 246, 135, 130, 96, 229, 38, 22, 235,
			32, 42, 193, 179, 58, 165, 215, 113, 160, 235, 64, 190, 83, 28,
			8, 218, 236, 50, 158, 214, 5, 45, 85, 178, 190, 164, 1, 23,
			192, 213, 146, 10, 40, 18, 69, 5, 75, 114, 121, 13, 181, 185,
			32, 181, 164, 8, 54, 163, 191, 74, 202, 167, 204, 236, 212, 142,
			203, 164, 160, 58, 160, 205, 218, 42, 249, 181, 206, 224, 206, 148,
			212, 25, 149, 141, 43, 46, 36, 20, 201, 10, 102, 84, 87, 74,
			6, 138, 3, 45, 51, 46, 36, 213, 69, 81, 9, 94, 112, 69,
			161, 193, 68, 73, 200, 168, 96, 23, 52, 131, 185, 224, 5, 106,
			80, 144, 124, 174, 46, 117, 153, 180, 21, 4, 178, 162, 169, 174,
			32, 168, 4, 211, 133, 37, 116, 237, 148, 77, 21, 73, 105, 98,
			71, 16, 143, 73, 4, 81, 112, 20, 159, 57, 33, 6, 18, 193,
			36, 12, 78, 137, 139, 93, 24, 158, 67, 60, 198, 48, 10, 38,
			231, 33, 57, 30, 199, 48, 14, 60, 23, 135, 17, 56, 190, 11,
			163, 192, 143, 67, 50, 156, 198, 65, 24, 33, 216, 119, 34, 32,
			209, 190, 209, 56, 254, 57, 224, 47, 38, 33, 142, 34, 8, 66,
			32, 39, 19, 143, 96, 23, 206, 156, 48, 116, 252, 152, 224, 168,
			15, 196, 31, 121, 83, 151, 248, 199, 125, 24, 78, 99, 240, 131,
			24, 129, 71, 78, 72, 140, 93, 136, 131, 190, 57, 246, 219, 251,
			32, 56, 130, 19, 28, 142, 198, 142, 31, 59, 67, 226, 145, 248,
			220, 28, 120, 68, 98, 95, 31, 118, 20, 132, 8, 28, 152, 56,
			97, 76, 70, 83, 207, 9, 97, 50, 13, 39, 65, 132, 65, 103,
			230, 146, 104, 228, 57, 228, 4, 187, 54, 16, 31, 252, 0, 240,
			41, 246, 99, 136, 198, 142, 231, 221, 76, 20, 65, 112, 230, 227,
			80, 71, 191, 153, 38, 12, 49, 120, 196, 25, 122, 88, 31, 101,
			242, 116, 73, 136, 71, 177, 78, 232, 122, 53, 34, 46, 246, 99,
			199, 235, 35, 136, 38, 120, 68, 28, 175, 15, 248, 11, 124, 50,
			241, 156, 240, 188, 223, 58, 141, 240, 111, 166, 216, 143, 137, 227,
			129, 235, 156, 56, 199, 56, 130, 103, 255, 11, 149, 73, 24, 140,
			166, 33, 62, 209, 81, 7, 71, 16, 77, 135, 81, 76, 226, 105,
			140, 225, 56, 8, 92, 3, 118, 132, 195, 83, 50, 194, 209, 27,
			240, 130, 200, 0, 54, 141, 112, 31, 129, 235, 196, 142, 57, 122,
			18, 6, 71, 36, 142, 222, 232, 245, 112, 26, 17, 3, 28, 241,
			99, 28, 134, 211, 73, 76, 2, 255, 57, 140, 131, 51, 124, 138,
			67, 24, 57, 211, 8, 187, 6, 225, 192, 215, 217, 234, 90, 193,
			65, 120, 174, 221, 106, 28, 204, 13, 244, 225, 108, 140, 227, 49,
			14, 53, 168, 6, 45, 71, 195, 16, 197, 33, 25, 197, 155, 102,
			65, 8, 113, 16, 198, 104, 35, 79, 240, 241, 177, 71, 142, 177,
			63, 194, 90, 29, 104, 55, 103, 36, 194, 207, 193, 9, 73, 164,
			13, 136, 57, 24, 206, 156, 115, 8, 166, 38, 107, 125, 81, 211,
			8, 163, 102, 189, 81, 186, 125, 115, 159, 64, 142, 192, 113, 79,
			137, 142, 188, 181, 158, 4, 81, 68, 218, 114, 49, 176, 141, 198,
			45, 230, 54, 66, 187, 104, 171, 211, 235, 194, 238, 158, 94, 237,
			246, 186, 251, 214, 27, 244, 61, 212, 217, 253, 172, 89, 54, 194,
			159, 90, 191, 54, 194, 123, 205, 178, 17, 62, 177, 250, 70, 184,
			213, 44, 27, 225, 103, 214, 47, 140, 176, 93, 54, 194, 159, 89,
			251, 70, 136, 154, 101, 35, 124, 106, 125, 106, 132, 79, 154, 101,
			35, 124, 102, 125, 98, 132, 159, 52, 203, 127, 117, 80, 103, 219,
			234, 117, 95, 90, 15, 30, 253, 163, 3, 14, 44, 104, 73, 5,
			75, 193, 240, 39, 20, 84, 202, 100, 65, 27, 10, 88, 241, 26,
			210, 164, 4, 65, 95, 104, 162, 81, 28, 146, 11, 206, 50, 200,
			232, 156, 149, 166, 253, 213, 85, 174, 201, 132, 102, 232, 230, 126,
			211, 126, 87, 188, 22, 224, 76, 136, 180, 193, 1, 181, 170, 88,
			154, 228, 64, 223, 39, 69, 149, 83, 96, 82, 251, 51, 252, 165,
			32, 145, 166, 139, 9, 250, 187, 154, 74, 133, 160, 237, 106, 130,
			202, 138, 151, 250, 228, 85, 101, 90, 95, 82, 106, 127, 154, 124,
			150, 60, 179, 225, 136, 11, 96, 165, 84, 73, 153, 210, 53, 27,
			105, 126, 101, 41, 133, 35, 206, 225, 235, 70, 4, 32, 170, 20,
			134, 137, 120, 118, 107, 200, 176, 205, 140, 241, 92, 115, 83, 45,
			74, 9, 31, 208, 191, 105, 220, 124, 163, 27, 219, 146, 194, 219,
			40, 240, 13, 147, 80, 121, 213, 230, 231, 92, 192, 59, 99, 253,
			78, 103, 214, 96, 97, 12, 249, 236, 183, 52, 85, 240, 238, 235,
			111, 222, 217, 8, 33, 212, 221, 182, 182, 122, 221, 151, 187, 223,
			159, 237, 152, 99, 94, 162, 191, 223, 69, 163, 5, 183, 211, 165,
			224, 5, 171, 11, 155, 139, 197, 32, 175, 83, 54, 104, 161, 146,
			131, 175, 94, 201, 193, 146, 106, 206, 225, 34, 207, 6, 73, 197,
			170, 217, 134, 160, 29, 177, 126, 162, 55, 217, 235, 77, 246, 87,
			175, 164, 125, 109, 243, 232, 191, 77, 103, 135, 46, 186, 123, 44,
			40, 85, 84, 244, 126, 133, 238, 68, 201, 106, 204, 122, 15, 111,
			143, 100, 13, 26, 143, 62, 32, 223, 183, 134, 191, 252, 242, 240,
			255, 207, 227, 237, 95, 239, 232, 57, 236, 35, 235, 71, 91, 232,
			47, 219, 102, 14, 251, 200, 234, 29, 254, 105, 251, 198, 72, 117,
			120, 96, 160, 247, 166, 35, 2, 78, 173, 150, 92, 72, 205, 51,
			30, 75, 105, 169, 137, 173, 46, 179, 150, 37, 157, 42, 73, 181,
			101, 163, 233, 195, 41, 21, 154, 149, 224, 208, 62, 128, 103, 218,
			96, 191, 85, 237, 235, 123, 213, 21, 174, 9, 178, 228, 202, 212,
			162, 225, 188, 57, 203, 41, 208, 247, 41, 173, 148, 46, 227, 148,
			23, 85, 206, 116, 141, 93, 177, 245, 218, 189, 141, 224, 188, 245,
			192, 103, 102, 190, 73, 204, 56, 161, 107, 117, 195, 12, 18, 213,
			86, 167, 153, 250, 94, 15, 6, 151, 151, 151, 118, 98, 34, 109,
			96, 106, 236, 228, 192, 35, 35, 236, 71, 248, 197, 161, 125, 128,
			16, 76, 203, 156, 74, 105, 158, 5, 19, 52, 131, 217, 10, 146,
			202, 188, 184, 89, 78, 33, 79, 46, 245, 67, 73, 22, 130, 54,
			212, 206, 74, 67, 199, 172, 92, 244, 175, 120, 123, 99, 174, 184,
			1, 211, 58, 50, 38, 111, 24, 152, 137, 229, 138, 121, 135, 78,
			68, 162, 62, 130, 51, 18, 143, 117, 171, 220, 164, 77, 195, 56,
			46, 209, 237, 221, 112, 130, 110, 169, 159, 19, 223, 237, 67, 59,
			178, 208, 247, 250, 133, 72, 29, 34, 211, 0, 154, 161, 55, 162,
			244, 198, 241, 243, 246, 165, 95, 77, 21, 122, 180, 175, 117, 3,
			90, 240, 11, 42, 76, 147, 185, 30, 45, 204, 4, 134, 32, 103,
			5, 107, 222, 157, 252, 118, 70, 87, 253, 247, 193, 238, 147, 182,
			3, 246, 44, 178, 110, 181, 237, 178, 107, 245, 186, 63, 188, 251,
			4, 33, 212, 217, 177, 122, 219, 15, 117, 241, 33, 212, 221, 209,
			79, 243, 225, 238, 199, 232, 30, 218, 222, 177, 58, 86, 175, 187,
			215, 193, 232, 62, 186, 163, 63, 182, 122, 221, 189, 157, 123, 235,
			175, 78, 175, 187, 119, 255, 211, 245, 87, 183, 215, 221, 235, 59,
			235, 23, 253, 159, 0, 0, 0, 255, 255, 92, 125, 20, 219, 13,
			13, 0, 0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptor.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("luci.examples.k8s.helloworld.Greeter")
	if err != nil {
		panic(err)
	}
	return ret
}