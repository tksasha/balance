# frozen_string_literal: true

RSpec.describe ApplicationHelper do
  subject { helper }

  describe '#new_item_for_inline_form' do
    context do
      before do
        allow(subject).to receive(:currency_from_params).and_return('usd')

        allow(Item).to receive(:new).with(currency: 'usd').and_return(:item)
      end

      its(:new_item_for_inline_form) { is_expected.to eq :item }
    end

    context do
      before do
        allow(subject).to receive(:currency_from_params).and_return('uah')

        allow(Item).to receive(:new).with(currency: 'uah').and_return(:item)
      end

      its(:new_item_for_inline_form) { is_expected.to eq :item }
    end
  end
end
