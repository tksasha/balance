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
            render 'admin/dashboard/cashes/currency', currency:, collection:
          end
        end
    end

    h2 'Витрати по категоріям'
    columns do
      Item
        .joins(:category)
        .where(items: { currency: CURRENCIES.keys })
        .group('items.currency, categories.supercategory')
        .pluck('items.currency, categories.supercategory, SUM(items.sum) AS sum')
        .group_by(&:first)
        .map do |currency, collection|
          column do
            h3 currency.upcase

            render 'admin/dashboard/items/currency', currency:, collection:
          end
        end
    end
  end
end
