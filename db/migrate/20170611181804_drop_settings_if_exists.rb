class DropSettingsIfExists < ActiveRecord::Migration[5.1]
  def change
    drop_table :settings if table_exists? :settings
  end
end
