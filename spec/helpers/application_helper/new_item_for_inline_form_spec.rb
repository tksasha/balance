# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  subject { helper }

  describe '#new_item_for_inline_form' do
    context do
      before { allow(subject).to receive(:currency_from_params).and_return('usd') }

      before { allow(Item).to receive(:new).with(currency: 'usd').and_return(:item) }

      its(:new_item_for_inline_form) { is_expected.to eq :item }
    end

    context do
      before { allow(subject).to receive(:currency_from_params).and_return('uah') }

      before { allow(Item).to receive(:new).with(currency: 'uah').and_return(:item) }

      its(:new_item_for_inline_form) { is_expected.to eq :item }
    end
  end
end
