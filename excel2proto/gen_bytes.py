import openpyxl, activity_weight_pb2, pathlib, struct

SRC = "conf.xlsx"
OUT = pathlib.Path("output/activity_weight.bytes")
OUT.parent.mkdir(exist_ok=True)

wb = openpyxl.load_workbook(SRC, data_only=True)
ws = wb.active  # Sheet1

cfg = activity_weight_pb2.ActivityWeightList()

for row in ws.iter_rows(min_row=2, values_only=True):
    item = cfg.items.add()
    item.Id, item.Weight = map(int, row)

OUT.write_bytes(cfg.SerializeToString())
print("✔ 导出完成 ->", OUT)