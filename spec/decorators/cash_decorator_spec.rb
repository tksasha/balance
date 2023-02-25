# frozen_string_literal: true

RSpec.describe CashDecorator do
  subject { cash.decorate }

  describe '#css_class' do
    let(:cash) { build(:cash, favorite:) }

    context 'when it is favorite?' do
      let(:favorite) { true }

      its(:css_class) { is_expected.to eq 'cash cash-favorite' }
    end

    context 'when it is not favorite?' do
      let(:favorite) { false }

      its(:css_class) { is_expected.to eq 'cash' }
    end
  end
end
