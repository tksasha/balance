class AddFavoriteToCashes < ActiveRecord::Migration[7.0]
  def change
    add_column :cashes, :favorite, :boolean, default: false
  end
end
