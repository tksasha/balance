# frozen_string_literal: true

ActiveAdmin.register_page 'Dashboard' do
  menu priority: 1, label: proc { I18n.t('active_admin.dashboard') }

  content title: proc { I18n.t('active_admin.dashboard') } do
    h2 'Залишки'
    columns do
      Cash
        .where(currency: CURRENCIES.keys)
        .group(:currency, :supercategory)
        .pluck('currency, supercategory, SUM(sum) AS sum')
        .group_by(&:first)
        .map do |currency, collection|
          column do
            render 'currency', currency:, collection:
          end
        end
    end
  end
end
