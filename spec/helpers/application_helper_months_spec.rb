# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  describe '#months' do
    subject { helper.months }

    context 'when locale is :ua' do
      it do
        I18n.with_locale(:ua) do
          expect(subject).to eq \
            %w[Січень Лютий Березень Квітень Травень Червень Липень Серпень Вересень Жовтень Листопад Грудень]
        end
      end
    end

    context 'when locale is :en' do
      it do
        I18n.with_locale(:en) do
          expect(subject).to eq \
            %w[January February March April May June July August September October November December]
        end
      end
    end
  end
end
