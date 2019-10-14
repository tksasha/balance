class ChangeColumnNullForSumInCashes < ActiveRecord::Migration[6.0]
  def change
    change_column_null :cashes, :sum, true 
  end
end
