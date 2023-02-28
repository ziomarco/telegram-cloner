package http

type GetChatsResponse struct {
	Ok     bool `json:"ok"`
	Result []struct {
		Id        int    `json:"id"`
		Type      string `json:"type"`
		Title     string `json:"title"`
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Photo     struct {
			SmallFileId       string `json:"small_file_id"`
			SmallFileUniqueId string `json:"small_file_unique_id"`
			BigFileId         string `json:"big_file_id"`
			BigFileUniqueId   string `json:"big_file_unique_id"`
		} `json:"photo"`
		Bio           string `json:"bio"`
		Description   string `json:"description"`
		InviteLink    string `json:"invite_link"`
		PinnedMessage struct {
			MessageId int `json:"message_id"`
			From      struct {
				Id                      int    `json:"id"`
				IsBot                   bool   `json:"is_bot"`
				FirstName               string `json:"first_name"`
				LastName                string `json:"last_name"`
				Username                string `json:"username"`
				LanguageCode            string `json:"language_code"`
				CanJoinGroups           bool   `json:"can_join_groups"`
				CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
				SupportsInlineQueries   bool   `json:"supports_inline_queries"`
				IsVerified              bool   `json:"is_verified"`
				IsScam                  bool   `json:"is_scam"`
			} `json:"from"`
			SenderChat  string `json:"sender_chat"`
			Date        int    `json:"date"`
			Chat        string `json:"chat"`
			ForwardFrom struct {
				Id                      int    `json:"id"`
				IsBot                   bool   `json:"is_bot"`
				FirstName               string `json:"first_name"`
				LastName                string `json:"last_name"`
				Username                string `json:"username"`
				LanguageCode            string `json:"language_code"`
				CanJoinGroups           bool   `json:"can_join_groups"`
				CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
				SupportsInlineQueries   bool   `json:"supports_inline_queries"`
				IsVerified              bool   `json:"is_verified"`
				IsScam                  bool   `json:"is_scam"`
			} `json:"forward_from"`
			ForwardFromChat      string `json:"forward_from_chat"`
			ForwardFromMessageId int    `json:"forward_from_message_id"`
			ForwardSignature     string `json:"forward_signature"`
			ForwardSenderName    string `json:"forward_sender_name"`
			ForwardDate          int    `json:"forward_date"`
			ReplyToMessage       string `json:"reply_to_message"`
			ViaBot               struct {
				Id                      int    `json:"id"`
				IsBot                   bool   `json:"is_bot"`
				FirstName               string `json:"first_name"`
				LastName                string `json:"last_name"`
				Username                string `json:"username"`
				LanguageCode            string `json:"language_code"`
				CanJoinGroups           bool   `json:"can_join_groups"`
				CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
				SupportsInlineQueries   bool   `json:"supports_inline_queries"`
				IsVerified              bool   `json:"is_verified"`
				IsScam                  bool   `json:"is_scam"`
			} `json:"via_bot"`
			EditDate        int    `json:"edit_date"`
			MediaGroupId    string `json:"media_group_id"`
			MessageThreadId string `json:"message_thread_id"`
			AuthorSignature string `json:"author_signature"`
			Text            string `json:"text"`
			Entities        []struct {
				Type   string `json:"type"`
				Offset int    `json:"offset"`
				Length int    `json:"length"`
				Url    string `json:"url"`
				User   struct {
					Id                      int    `json:"id"`
					IsBot                   bool   `json:"is_bot"`
					FirstName               string `json:"first_name"`
					LastName                string `json:"last_name"`
					Username                string `json:"username"`
					LanguageCode            string `json:"language_code"`
					CanJoinGroups           bool   `json:"can_join_groups"`
					CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
					SupportsInlineQueries   bool   `json:"supports_inline_queries"`
					IsVerified              bool   `json:"is_verified"`
					IsScam                  bool   `json:"is_scam"`
				} `json:"user"`
				Language string `json:"language"`
			} `json:"entities"`
			Animation struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				Duration     int    `json:"duration"`
				Thumb        struct {
					FileId       string `json:"file_id"`
					FileUniqueId string `json:"file_unique_id"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					FileSize     int    `json:"file_size"`
				} `json:"thumb"`
				FileName string `json:"file_name"`
				MimeType string `json:"mime_type"`
				FileSize int    `json:"file_size"`
			} `json:"animation"`
			Audio struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				Duration     int    `json:"duration"`
				Performer    string `json:"performer"`
				Title        string `json:"title"`
				FileName     string `json:"file_name"`
				MimeType     string `json:"mime_type"`
				FileSize     int    `json:"file_size"`
				Thumb        struct {
					FileId       string `json:"file_id"`
					FileUniqueId string `json:"file_unique_id"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					FileSize     int    `json:"file_size"`
				} `json:"thumb"`
			} `json:"audio"`
			Document struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				Thumb        struct {
					FileId       string `json:"file_id"`
					FileUniqueId string `json:"file_unique_id"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					FileSize     int    `json:"file_size"`
				} `json:"thumb"`
				FileName string `json:"file_name"`
				MimeType string `json:"mime_type"`
				FileSize int    `json:"file_size"`
			} `json:"document"`
			Photo []struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				FileSize     int    `json:"file_size"`
			} `json:"photo"`
			Sticker struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				IsAnimated   bool   `json:"is_animated"`
				Thumb        struct {
					FileId       string `json:"file_id"`
					FileUniqueId string `json:"file_unique_id"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					FileSize     int    `json:"file_size"`
				} `json:"thumb"`
				Emoji        string `json:"emoji"`
				SetName      string `json:"set_name"`
				MaskPosition struct {
					Point  string `json:"point"`
					XShift int    `json:"x_shift"`
					YShift int    `json:"y_shift"`
					Scale  int    `json:"scale"`
				} `json:"mask_position"`
				FileSize int `json:"file_size"`
			} `json:"sticker"`
			Video struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				Duration     int    `json:"duration"`
				Thumb        struct {
					FileId       string `json:"file_id"`
					FileUniqueId string `json:"file_unique_id"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					FileSize     int    `json:"file_size"`
				} `json:"thumb"`
				FileName string `json:"file_name"`
				MimeType string `json:"mime_type"`
				FileSize int    `json:"file_size"`
			} `json:"video"`
			VideoNote struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				Length       int    `json:"length"`
				Duration     int    `json:"duration"`
				Thumb        struct {
					FileId       string `json:"file_id"`
					FileUniqueId string `json:"file_unique_id"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					FileSize     int    `json:"file_size"`
				} `json:"thumb"`
				FileSize int `json:"file_size"`
			} `json:"video_note"`
			Voice struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				Duration     int    `json:"duration"`
				MimeType     string `json:"mime_type"`
				FileSize     int    `json:"file_size"`
			} `json:"voice"`
			Caption         string `json:"caption"`
			CaptionEntities []struct {
				Type   string `json:"type"`
				Offset int    `json:"offset"`
				Length int    `json:"length"`
				Url    string `json:"url"`
				User   struct {
					Id                      int    `json:"id"`
					IsBot                   bool   `json:"is_bot"`
					FirstName               string `json:"first_name"`
					LastName                string `json:"last_name"`
					Username                string `json:"username"`
					LanguageCode            string `json:"language_code"`
					CanJoinGroups           bool   `json:"can_join_groups"`
					CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
					SupportsInlineQueries   bool   `json:"supports_inline_queries"`
					IsVerified              bool   `json:"is_verified"`
					IsScam                  bool   `json:"is_scam"`
				} `json:"user"`
				Language string `json:"language"`
			} `json:"caption_entities"`
			Contact struct {
				PhoneNumber string `json:"phone_number"`
				FirstName   string `json:"first_name"`
				LastName    string `json:"last_name"`
				UserId      int    `json:"user_id"`
				Vcard       string `json:"vcard"`
			} `json:"contact"`
			Dice struct {
				Emoji string `json:"emoji"`
				Value int    `json:"value"`
			} `json:"dice"`
			Game struct {
				Title       string `json:"title"`
				Description string `json:"description"`
				Photo       []struct {
					FileId       string `json:"file_id"`
					FileUniqueId string `json:"file_unique_id"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					FileSize     int    `json:"file_size"`
				} `json:"photo"`
				Text         string `json:"text"`
				TextEntities []struct {
					Type   string `json:"type"`
					Offset int    `json:"offset"`
					Length int    `json:"length"`
					Url    string `json:"url"`
					User   struct {
						Id                      int    `json:"id"`
						IsBot                   bool   `json:"is_bot"`
						FirstName               string `json:"first_name"`
						LastName                string `json:"last_name"`
						Username                string `json:"username"`
						LanguageCode            string `json:"language_code"`
						CanJoinGroups           bool   `json:"can_join_groups"`
						CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
						SupportsInlineQueries   bool   `json:"supports_inline_queries"`
						IsVerified              bool   `json:"is_verified"`
						IsScam                  bool   `json:"is_scam"`
					} `json:"user"`
					Language string `json:"language"`
				} `json:"text_entities"`
				Animation struct {
					FileId       string `json:"file_id"`
					FileUniqueId string `json:"file_unique_id"`
					Width        int    `json:"width"`
					Height       int    `json:"height"`
					Duration     int    `json:"duration"`
					Thumb        struct {
						FileId       string `json:"file_id"`
						FileUniqueId string `json:"file_unique_id"`
						Width        int    `json:"width"`
						Height       int    `json:"height"`
						FileSize     int    `json:"file_size"`
					} `json:"thumb"`
					FileName string `json:"file_name"`
					MimeType string `json:"mime_type"`
					FileSize int    `json:"file_size"`
				} `json:"animation"`
			} `json:"game"`
			Poll struct {
				Id       string `json:"id"`
				Question string `json:"question"`
				Options  []struct {
					Text       string `json:"text"`
					VoterCount int    `json:"voter_count"`
				} `json:"options"`
				TotalVoterCount       int    `json:"total_voter_count"`
				IsClosed              bool   `json:"is_closed"`
				IsAnonymous           bool   `json:"is_anonymous"`
				Type                  string `json:"type"`
				AllowsMultipleAnswers bool   `json:"allows_multiple_answers"`
				CorrectOptionId       int    `json:"correct_option_id"`
				Explanation           string `json:"explanation"`
				ExplanationEntities   []struct {
					Type   string `json:"type"`
					Offset int    `json:"offset"`
					Length int    `json:"length"`
					Url    string `json:"url"`
					User   struct {
						Id                      int    `json:"id"`
						IsBot                   bool   `json:"is_bot"`
						FirstName               string `json:"first_name"`
						LastName                string `json:"last_name"`
						Username                string `json:"username"`
						LanguageCode            string `json:"language_code"`
						CanJoinGroups           bool   `json:"can_join_groups"`
						CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
						SupportsInlineQueries   bool   `json:"supports_inline_queries"`
						IsVerified              bool   `json:"is_verified"`
						IsScam                  bool   `json:"is_scam"`
					} `json:"user"`
					Language string `json:"language"`
				} `json:"explanation_entities"`
				OpenPeriod int `json:"open_period"`
				CloseDate  int `json:"close_date"`
			} `json:"poll"`
			Venue struct {
				Location struct {
					Longitude            int `json:"longitude"`
					Latitude             int `json:"latitude"`
					HorizontalAccuracy   int `json:"horizontal_accuracy"`
					LivePeriod           int `json:"live_period"`
					Heading              int `json:"heading"`
					ProximityAlertRadius int `json:"proximity_alert_radius"`
				} `json:"location"`
				Title           string `json:"title"`
				Address         string `json:"address"`
				FoursquareId    string `json:"foursquare_id"`
				FoursquareType  string `json:"foursquare_type"`
				GooglePlaceId   string `json:"google_place_id"`
				GooglePlaceType string `json:"google_place_type"`
			} `json:"venue"`
			Location struct {
				Longitude            int `json:"longitude"`
				Latitude             int `json:"latitude"`
				HorizontalAccuracy   int `json:"horizontal_accuracy"`
				LivePeriod           int `json:"live_period"`
				Heading              int `json:"heading"`
				ProximityAlertRadius int `json:"proximity_alert_radius"`
			} `json:"location"`
			NewChatMembers []struct {
				Id                      int    `json:"id"`
				IsBot                   bool   `json:"is_bot"`
				FirstName               string `json:"first_name"`
				LastName                string `json:"last_name"`
				Username                string `json:"username"`
				LanguageCode            string `json:"language_code"`
				CanJoinGroups           bool   `json:"can_join_groups"`
				CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
				SupportsInlineQueries   bool   `json:"supports_inline_queries"`
				IsVerified              bool   `json:"is_verified"`
				IsScam                  bool   `json:"is_scam"`
			} `json:"new_chat_members"`
			LeftChatMember struct {
				Id                      int    `json:"id"`
				IsBot                   bool   `json:"is_bot"`
				FirstName               string `json:"first_name"`
				LastName                string `json:"last_name"`
				Username                string `json:"username"`
				LanguageCode            string `json:"language_code"`
				CanJoinGroups           bool   `json:"can_join_groups"`
				CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
				SupportsInlineQueries   bool   `json:"supports_inline_queries"`
				IsVerified              bool   `json:"is_verified"`
				IsScam                  bool   `json:"is_scam"`
			} `json:"left_chat_member"`
			NewChatTitle string `json:"new_chat_title"`
			NewChatPhoto []struct {
				FileId       string `json:"file_id"`
				FileUniqueId string `json:"file_unique_id"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
				FileSize     int    `json:"file_size"`
			} `json:"new_chat_photo"`
			DeleteChatPhoto       bool   `json:"delete_chat_photo"`
			GroupChatCreated      bool   `json:"group_chat_created"`
			SupergroupChatCreated bool   `json:"supergroup_chat_created"`
			ChannelChatCreated    bool   `json:"channel_chat_created"`
			MigrateToChatId       int    `json:"migrate_to_chat_id"`
			MigrateFromChatId     int    `json:"migrate_from_chat_id"`
			PinnedMessage         string `json:"pinned_message"`
			Invoice               struct {
				Title          string `json:"title"`
				Description    string `json:"description"`
				StartParameter string `json:"start_parameter"`
				Currency       string `json:"currency"`
				TotalAmount    int    `json:"total_amount"`
			} `json:"invoice"`
			SuccessfulPayment struct {
				Currency         string `json:"currency"`
				TotalAmount      int    `json:"total_amount"`
				InvoicePayload   string `json:"invoice_payload"`
				ShippingOptionId string `json:"shipping_option_id"`
				OrderInfo        struct {
					Name            string `json:"name"`
					PhoneNumber     string `json:"phone_number"`
					Email           string `json:"email"`
					ShippingAddress struct {
						CountryCode string `json:"country_code"`
						State       string `json:"state"`
						City        string `json:"city"`
						StreetLine1 string `json:"street_line1"`
						StreetLine2 string `json:"street_line2"`
						PostCode    string `json:"post_code"`
					} `json:"shipping_address"`
				} `json:"order_info"`
				TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
				ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
			} `json:"successful_payment"`
			ConnectedWebsite string `json:"connected_website"`
			PassportData     struct {
				Data []struct {
					Type        string `json:"type"`
					Data        string `json:"data"`
					PhoneNumber string `json:"phone_number"`
					Email       string `json:"email"`
					Files       []struct {
						FileId       string `json:"file_id"`
						FileUniqueId string `json:"file_unique_id"`
						FileSize     int    `json:"file_size"`
						FileDate     int    `json:"file_date"`
					} `json:"files"`
					FrontSide struct {
						FileId       string `json:"file_id"`
						FileUniqueId string `json:"file_unique_id"`
						FileSize     int    `json:"file_size"`
						FileDate     int    `json:"file_date"`
					} `json:"front_side"`
					ReverseSide struct {
						FileId       string `json:"file_id"`
						FileUniqueId string `json:"file_unique_id"`
						FileSize     int    `json:"file_size"`
						FileDate     int    `json:"file_date"`
					} `json:"reverse_side"`
					Selfie struct {
						FileId       string `json:"file_id"`
						FileUniqueId string `json:"file_unique_id"`
						FileSize     int    `json:"file_size"`
						FileDate     int    `json:"file_date"`
					} `json:"selfie"`
					Translation []struct {
						FileId       string `json:"file_id"`
						FileUniqueId string `json:"file_unique_id"`
						FileSize     int    `json:"file_size"`
						FileDate     int    `json:"file_date"`
					} `json:"translation"`
					Hash string `json:"hash"`
				} `json:"data"`
				Credentials struct {
					Data   string `json:"data"`
					Hash   string `json:"hash"`
					Secret string `json:"secret"`
				} `json:"credentials"`
			} `json:"passport_data"`
			ProximityAlertTriggered struct {
				Traveler struct {
					Id                      int    `json:"id"`
					IsBot                   bool   `json:"is_bot"`
					FirstName               string `json:"first_name"`
					LastName                string `json:"last_name"`
					Username                string `json:"username"`
					LanguageCode            string `json:"language_code"`
					CanJoinGroups           bool   `json:"can_join_groups"`
					CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
					SupportsInlineQueries   bool   `json:"supports_inline_queries"`
					IsVerified              bool   `json:"is_verified"`
					IsScam                  bool   `json:"is_scam"`
				} `json:"traveler"`
				Watcher struct {
					Id                      int    `json:"id"`
					IsBot                   bool   `json:"is_bot"`
					FirstName               string `json:"first_name"`
					LastName                string `json:"last_name"`
					Username                string `json:"username"`
					LanguageCode            string `json:"language_code"`
					CanJoinGroups           bool   `json:"can_join_groups"`
					CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
					SupportsInlineQueries   bool   `json:"supports_inline_queries"`
					IsVerified              bool   `json:"is_verified"`
					IsScam                  bool   `json:"is_scam"`
				} `json:"watcher"`
				Distance int `json:"distance"`
			} `json:"proximity_alert_triggered"`
			ReplyMarkup struct {
				InlineKeyboard [][]struct {
					Text     string `json:"text"`
					Url      string `json:"url"`
					LoginUrl struct {
						Url                string `json:"url"`
						ForwardText        string `json:"forward_text"`
						BotUsername        string `json:"bot_username"`
						RequestWriteAccess bool   `json:"request_write_access"`
					} `json:"login_url"`
					CallbackData                 string `json:"callback_data"`
					SwitchInlineQuery            string `json:"switch_inline_query"`
					SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat"`
					CallbackGame                 string `json:"callback_game"`
					Pay                          bool   `json:"pay"`
				} `json:"inline_keyboard"`
			} `json:"reply_markup"`
			Views       int    `json:"views"`
			Forwards    int    `json:"forwards"`
			IsScheduled bool   `json:"is_scheduled"`
			ScheduledAt string `json:"scheduled_at"`
		} `json:"pinned_message"`
		Permissions struct {
			CanSendMessages       bool `json:"can_send_messages"`
			CanSendMediaMessages  bool `json:"can_send_media_messages"`
			CanSendPolls          bool `json:"can_send_polls"`
			CanSendOtherMessages  bool `json:"can_send_other_messages"`
			CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
			CanChangeInfo         bool `json:"can_change_info"`
			CanInviteUsers        bool `json:"can_invite_users"`
			CanPinMessages        bool `json:"can_pin_messages"`
		} `json:"permissions"`
		SlowModeDelay    int    `json:"slow_mode_delay"`
		StickerSetName   string `json:"sticker_set_name"`
		CanSetStickerSet bool   `json:"can_set_sticker_set"`
		LinkedChatId     int    `json:"linked_chat_id"`
		Location         struct {
			Location struct {
				Longitude            int `json:"longitude"`
				Latitude             int `json:"latitude"`
				HorizontalAccuracy   int `json:"horizontal_accuracy"`
				LivePeriod           int `json:"live_period"`
				Heading              int `json:"heading"`
				ProximityAlertRadius int `json:"proximity_alert_radius"`
			} `json:"location"`
			Address string `json:"address"`
		} `json:"location"`
		IsVerified bool `json:"is_verified"`
		IsScam     bool `json:"is_scam"`
		Distance   int  `json:"distance"`
	} `json:"result"`
}
