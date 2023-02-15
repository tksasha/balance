# frozen_string_literal: true

ActiveAdmin.register_page 'Dashboard' do
  menu priority: 1, label: proc { I18n.t('active_admin.dashboard') }

  content title: proc { I18n.t('active_admin.dashboard') } do
    h2 'Залишки'
    columns do
      Cash
        .group(:currency, :supercategory)
        .select('currency, supercategory, SUM(sum)')
        .pluck(:currency, :supercategory, :sum)
        .group_by(&:first)
        .map do |currency, collection|
          column do
            render 'currency', currency:, collection:
          end
        end
    end
  end
end
