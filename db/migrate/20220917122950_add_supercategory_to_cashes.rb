class AddSupercategoryToCashes < ActiveRecord::Migration[7.0]
  def change
    add_column :cashes, :supercategory, :integer, null: false, default: 1
  end
end
